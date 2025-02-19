package server

import (
	"github.com/bytedance/sonic"
	"github.com/osamikoyo/IM-wharehouse/internal/data"
	"github.com/osamikoyo/IM-wharehouse/internal/data/models"
	"github.com/osamikoyo/IM-wharehouse/pkg/config"
	"github.com/osamikoyo/IM-wharehouse/pkg/loger/loger"
	"github.com/streadway/amqp"
)

type RpcServer struct {
	Channel *amqp.Channel
	Data *data.Storage
	cfg *config.Config
	loger loger.Logger
}

func New(cfg *config.Config) (*RpcServer, error){
	storage, err := data.New(cfg)
	if err != nil {
		return nil, err
	}

	conn, err := amqp.Dial(cfg.AmqpUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RpcServer{
		Channel: ch,
		Data: storage,
		cfg: cfg,
		loger: loger.New(),
	}, nil
}

func (r *RpcServer) Run() error {
	msgs, err := r.Channel.Consume(
			r.cfg.RpcQueueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
	if err != nil {
		return err
	}

	var req models.Request

	for msg := range msgs {
		if err = sonic.Unmarshal(msg.Body, &req);err != nil{
			r.loger.Error().Err(err)
		}

		if req.Rezerv {
			err = r.Data.RezervProduct(req.Name)
			if err != nil{
				r.loger.Error().Err(err)
			}
		}

		err = r.Channel.Publish(
			"",
			msg.ReplyTo,
			false,
			false,
			amqp.Publishing{
				CorrelationId: msg.CorrelationId,
				Body: []byte("succes"),
			},
		)

		if err != nil{
			r.loger.Error().Err(err)
		}
	}
	return nil
}