package services

import (
	"github.com/streadway/amqp"
	"log"
	"testing"
)

const url = "amqp://ricnsmart:9ef16689fdaf@dev.ricnsmart.com:5672/"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func TestSend(t *testing.T) {
	c := NewRabbitMQConnection(url)
	err := c.Open()
	failOnError(err, "Failed to open a connection")
	ch, err := c.Channel()
	failOnError(err, "failed to open a channel")
	//defer ch.Close()
	//forever := make(chan bool)
	newSender(ch, "test_queue1", "test_queue1")
	newSender(ch, "test_queue2", "test_queue2")
	newSender(ch, "test_queue3", "test_queue3")
	newSender(ch, "test_queue4", "test_queue4")
	newSender(ch, "test_queue5", "test_queue5")
	newSender(ch, "test_queue6", "test_queue6")
	newSender(ch, "test_queue7", "test_queue7")
	newSender(ch, "test_queue8", "test_queue8")
	newSender(ch, "test_queue9", "test_queue9")
	//<-forever
}

func TestReceive(t *testing.T) {
	c := NewRabbitMQConnection(url)
	err := c.Open()
	failOnError(err, "Failed to open a connection")
	forever := make(chan bool)
	go newReceiver(c, "test_queue1")
	go newReceiver(c, "test_queue2")
	go newReceiver(c, "test_queue3")
	go newReceiver(c, "test_queue4")
	go newReceiver(c, "test_queue5")
	go newReceiver(c, "test_queue6")
	go newReceiver(c, "test_queue7")
	go newReceiver(c, "test_queue8")
	go newReceiver(c, "test_queue9")
	<-forever
}

func newSender(ch *amqp.Channel, queue, msg string) {

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

	body := msg
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf(" [x] Sent %s", body)
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
			log.Printf("%v Received a message: %s", queue, d.Body)
		}
	}()

	log.Printf(" %v waiting for messages.", queue)
	<-forever
}
