package queues

import (
	"log"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)


type MQConnection struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewConnection(url string) (*MQConnection, error) {
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error

	// Retry logic
for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(url)
		if err != nil {
			log.Printf("RabbitMQ connection failed: %v. Retrying...", err)
			time.Sleep(2 * time.Second)
			continue
		}

		ch, err = conn.Channel()
		if err != nil {
			log.Printf("Failed to open channel: %v. Retrying...", err)
			conn.Close() // be sure to close bad connection
			time.Sleep(2 * time.Second)
			continue
		}

		// Success
		return &MQConnection{Conn: conn, Channel: ch}, nil
	}

	// If we reach here, all retries failed
	return nil, err
}

func (r *MQConnection) Close() {
	r.Channel.Close()
	r.Conn.Close()
}