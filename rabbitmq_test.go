package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"testing"
	"time"
)

const rabbitMQURL = "amqp://"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func TestSend(t *testing.T) {
	c := NewRabbitMQConnection(rabbitMQURL)
	err := c.Open()
	failOnError(err, "Failed to open a connection")
	ch, err := c.Channel()
	failOnError(err, "failed to open a channel")
	//defer ch.Close()
	//forever := make(chan bool)
	m := make(map[string]interface{})
	m["time"] = time.Now()
	newSender(ch, "test_queue1", m)
	//<-forever
}

func TestReceive(t *testing.T) {
	serviceName := "service"

	// 初始化zap日志
	InitZap(fmt.Sprintf(`config/%v.log`, serviceName), zap.String("service", serviceName))

	c := NewRabbitMQConnection(rabbitMQURL)
	err := c.Open()
	failOnError(err, "Failed to open a connection")
	forever := make(chan bool)
	go newReceiver(c, "test_queue1")
	//go newReceiver(c, "test_queue2")
	//go newReceiver(c, "test_queue3")
	//go newReceiver(c, "test_queue4")
	//go newReceiver(c, "test_queue5")
	//go newReceiver(c, "test_queue6")
	//go newReceiver(c, "test_queue7")
	//go newReceiver(c, "test_queue8")
	//go newReceiver(c, "test_queue9")
	<-forever
}

func newSender(ch *amqp.Channel, queue string, msg interface{}) {
	buf, _ := json.Marshal(msg)
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
		return
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        buf,
		})
	if err != nil {
		log.Println(err)
		return
	}
}

func newReceiver(conn *RabbitMQConnection, queue string) {
	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println(err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			m := make(map[string]interface{})
			err := json.Unmarshal(d.Body, &m)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(m["time"])
			t, err := time.Parse(time.RFC3339, m["time"].(string))
			if err != nil {
				log.Fatal(err)
			}
			m["time"] = t
			log.Println(t.String())
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			_, err = Collection("test").InsertOne(ctx, m)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Printf(" %v waiting for messages.", queue)
	<-forever
}
