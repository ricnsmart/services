package services

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
	"sync"
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
	mu         sync.Mutex       // 保护资源并发读写
	state      int32
}

func New(url string) *RabbitMQConnection {
	return &RabbitMQConnection{
		url:     url,
		stopCh:  make(chan struct{}),
		closeCh: make(chan *amqp.Error, 1),
	}
}

func (c *RabbitMQConnection) Open() error {
	// 进行Open期间不允许做任何跟连接有关的事情
	c.mu.Lock()
	defer c.mu.Unlock()

	conn, err := amqp.Dial(c.url)
	if err != nil {
		return err
	}

	atomic.StoreInt32(&c.state, OpenedState)
	c.connection = conn
	// 对connection的close事件加监听器
	c.connection.NotifyClose(c.closeCh)
	go c.keepAlive()

	return nil
}

func (c *RabbitMQConnection) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(c.stopCh)
}

func (c *RabbitMQConnection) keepAlive() {
	select {
	case <-c.stopCh:
		c.mu.Lock()
		c.connection.Close()
		atomic.StoreInt32(&c.state, ClosedState)
		c.mu.Unlock()
	case err := <-c.closeCh:
		log.Printf("rabbitMQ: disconnected with MQ, code:%d, reason:%s\n", err.Code, err.Reason)

		atomic.StoreInt32(&c.state, ReopeningState)
		var tempDelay time.Duration // how long to sleep on accept failure
		for {
			if e := c.Open(); e != nil {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("rabbitMQ: connection recover failed error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			log.Println("rabbitMQ: connection recover succeeded")
			return
		}
	}
}

func (c *RabbitMQConnection) Channel() (*amqp.Channel, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.State() != OpenedState {
		return nil, errors.New("rabbitMQ connection not opened")
	}
	return c.connection.Channel()
}

func (c *RabbitMQConnection) State() int32 {
	return atomic.LoadInt32(&c.state)
}
