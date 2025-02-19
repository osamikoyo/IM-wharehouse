package client

import (
	"github.com/bytedance/sonic"
	"github.com/osamikoyo/IM-wharehouse/pkg/config"
	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
)


type Sender struct{
	AmqpChannel *amqp.Channel
	AmqpQue amqp.Queue
}

func Init(cfg *config.Config) (*Sender, error) {
	conn, err := amqp.Dial(cfg.AmqpUrl)
	if err != nil{
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		return nil, err
	}
	defer ch.Close()

	que, err := ch.QueueDeclare(
		"message",
		false,
		false,
		false,
		false,
		nil,
	)

	return &Sender{
		AmqpQue: que,
		AmqpChannel: ch,
	}, err
}

func (s *Sender) Send(message *gomail.Message) error {
	body, err := sonic.Marshal(message)
	if err != nil{
		return err
	}

	err = s.AmqpChannel.Publish(
  		"",
  		s.AmqpQue.Name,
  		false,
  		false,
  		amqp.Publishing {
    		ContentType: "application/json",
    		Body:        []byte(body),
  	})

	if err != nil{
		return err
	}

	return nil
}