package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQManager struct {
	conn 			*amqp.Connection
	channel 		*amqp.Channel
	exchange 		string
}

func NewRabbitMQManager(url, exchange string)(*RabbitMQManager, error){
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("Error al conectarse con el servidor AMQP:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("Error al establecer canal:", err)
	}

	err = ch.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		conn.Close()
		ch.Close()
		return nil, fmt.Errorf("Error al establecer exchange:", err)
	}

	return &RabbitMQManager{
		conn: conn,
		channel: ch,
		exchange: exchange,
	}, nil
}

func (rbmqm *RabbitMQManager) PublishMessage(routingKey string, message interface{}) error {
	fmt.Println("Publicando mensaje en el tópico", routingKey, ":", message)

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("Error al convertir mensaje a JSON:", err)
	}

	err = rbmqm.channel.Publish(
		rbmqm.exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)

	if err != nil {
		return fmt.Errorf("Error al enviar mensaje por AMQP:", err)
	}

	fmt.Println("Mensaje publicado!")
	return nil
}

func (rbmqm *RabbitMQManager) CloseConnection() {
	rbmqm.channel.Close()
	rbmqm.conn.Close()
}