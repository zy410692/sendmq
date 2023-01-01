package AppInit

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

var MQConn *amqp.Connection

func GetConn() *amqp.Connection {
	return MQConn
}

func init() {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "admin", "123456", "192.168.31.227", 5672)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}
	MQConn = conn
	log.Println(MQConn.Major)
}
