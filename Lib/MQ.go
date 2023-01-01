package Lib

import (
	"log"
	"sendmq/AppInit"
	"strings"

	"github.com/streadway/amqp"
)

const (
	QUEUE_NEWUSER       = "newuser"
	QUEUE_NEWUSER_UNION = "newuser_union"
	EXCHANGE_USER       = "UserExchange"
	ROUTER_KEY_USERREG  = "userreg"
)

type MQ struct {
	Channel *amqp.Channel
}

func NewMq() *MQ {
	c, err := AppInit.GetConn().Channel()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &MQ{Channel: c}
}

func (this *MQ) DecQueueAndBind(queues string, key string, exchange string) error {
	qList := strings.Split(queues, ",")
	for _, queue := range qList {
		q, err := this.Channel.QueueDeclare(queue, false, false, false, false, nil)
		if err != nil {
			return err
		}
		err = this.Channel.QueueBind(q.Name, key, exchange, false, nil)
		if err != nil {
			return err
		}

	}
	return nil

}

func (this *MQ) SendMessage(key string, exchange string, message []byte) error {

	return this.Channel.Publish(exchange, key, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
}

func (this *MQ) Consume(queue string, key string, callback func(<-chan amqp.Delivery)) {
	msgs, err := this.Channel.Consume(queue, key, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	callback(msgs)
}
