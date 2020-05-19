package services

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
	"sync/atomic"
	"time"
)

var (
	ClosedState    = int32(0)
	OpenedState    = int32(1)
	ReopeningState = int32(2)
)

type RabbitMQConnection struct {
	url        string
	connection *amqp.Connection
	stopCh     chan struct{}
	closeCh    chan *amqp.Error // RabbitMQ 监听连接错误
	state      int32
}

func NewRabbitMQConnection(url string) *RabbitMQConnection {
	return &RabbitMQConnection{
		url:   url,
		state: ClosedState,
	}
}

func (c *RabbitMQConnection) Open() error {
	if c.State() == OpenedState {
		return errors.New("rabbitMQ: connection had been opened")
	}

	conn, err := amqp.Dial(c.url)
	if err != nil {
		return err
	}

	atomic.StoreInt32(&c.state, OpenedState)
	c.connection = conn
	c.stopCh = make(chan struct{})
	// 必须新建一个错误监听器，否则可能会导致无限关闭信号，导致反复重新open
	c.closeCh = make(chan *amqp.Error, 1)
	// 对connection的close事件加监听器
	c.connection.NotifyClose(c.closeCh)
	go c.keepAlive()

	return nil
}

func (c *RabbitMQConnection) Close() {
	if c.State() == ClosedState {
		return
	}
	select {
	case <-c.stopCh:
		// had been closed
	default:
		close(c.stopCh)
	}
	// wait done
	for c.State() != ClosedState {
		time.Sleep(time.Second)
	}
}

func (c *RabbitMQConnection) keepAlive() {
	select {
	case <-c.stopCh:
		c.connection.Close()
		atomic.StoreInt32(&c.state, ClosedState)
	case err := <-c.closeCh:
		if err != nil {
			log.Printf("disconnected with rabbitMQ,reason: %v", err)
		}

		atomic.StoreInt32(&c.state, ReopeningState)
		var tempDelay time.Duration // how long to sleep on accept failure
		for {
			select {
			case <-c.stopCh:
				c.connection.Close()
				atomic.StoreInt32(&c.state, ClosedState)
				return
			default:
				if e := c.Open(); e != nil {
					if tempDelay == 0 {
						tempDelay = 5 * time.Millisecond
					} else {
						tempDelay *= 2
					}
					if max := 1 * time.Second; tempDelay > max {
						tempDelay = max
					}

					log.Printf("rabbitMQ connection recover failed: %v,retry after %v", err, tempDelay)

					time.Sleep(tempDelay)
					continue
				}
				log.Printf("rabbitMQ connection recover succeeded")
				return
			}
		}
	}
}

// !! 当连接失败时，获取channel会出现阻塞，直到连接正常，才会返回可用的channel
func (c *RabbitMQConnection) Channel() (*amqp.Channel, error) {
	for c.State() != OpenedState {
		_, ok := <-c.stopCh
		if !ok {
			// 如果stopCh被主动关闭，结束Channel()
			return nil, errors.New("rabbitMQ connection had been closed")
		}
		// 否则继续等待，直到connection opened
		time.Sleep(time.Second)
	}
	return c.connection.Channel()
}

func (c *RabbitMQConnection) State() int32 {
	return atomic.LoadInt32(&c.state)
}
