package queues
import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewConsumer(ch *amqp.Channel, queueName string) (*Consumer, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		channel: ch,
		queue:   q,
	}, nil
}

func (c *Consumer) Consume(handler func(msg []byte)) error {
	msgs, err := c.channel.Consume(
		c.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Printf("Consuming on queue: %s", c.queue.Name)
	return nil
}