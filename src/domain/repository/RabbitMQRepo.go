package repository

type RabbitMQRepo interface {
	PublishMessage(routingKey string, message interface{}) error
	CloseConnection()
}