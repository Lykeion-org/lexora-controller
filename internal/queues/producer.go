package queues

import (
	"log"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)


type MQProducer struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewProducer(ch *amqp.Channel, queueName string) (*MQProducer, error) {
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

	return &MQProducer{
		channel: ch,
		queue:   q,
	}, nil
}

func (p *MQProducer) Publish(ctx context.Context, body []byte) error {
	err := p.channel.PublishWithContext(ctx,
		"",           // exchange
		p.queue.Name, // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Printf("Failed to publish: %v", err)
	}
	return err
}


