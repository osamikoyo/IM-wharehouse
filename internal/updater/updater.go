package updater

import (
	"github.com/dariubs/percent"
	"github.com/osamikoyo/IM-wharehouse/internal/data/models"
	"github.com/osamikoyo/IM-wharehouse/pkg/config"
	"github.com/osamikoyo/IM-wharehouse/pkg/loger/loger"
	"github.com/osamikoyo/IM-wharehouse/pkg/rpc/client"
)

type Updater struct {
	sender *client.Sender
	loger loger.Logger
}

func New(cfg *config.Config) (*Updater, error){
	sender, err := client.Init(cfg)
	if err != nil {
		return nil, err
	}

	return &Updater{
		sender: sender,
		loger: loger.New(),
	}, nil
}

const SMALLPROCENT = 10

func (u *Updater) Do(count, fullcount uint) {
	pr := percent.PercentOf(int(count), int(fullcount))
	if pr < SMALLPROCENT {
		msg := models.NewMessage(
			"example",
				"orgname",
				[]string{"example", ""},
				"products",
				"to small products number")
		err := u.sender.Send(msg.Message)
		if err != nil{
			u.loger.Error().Err(err)
		}
	}
}