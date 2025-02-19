package app

import (
	"fmt"

	"github.com/osamikoyo/IM-wharehouse/internal/rpc/server"
	"github.com/osamikoyo/IM-wharehouse/pkg/config"
	"github.com/osamikoyo/IM-wharehouse/pkg/loger/loger"
)

type App struct{
	loger loger.Logger
	rpcServer *server.RpcServer
	cfg *config.Config
}

func New() (*App, error){
	cfg, err := config.Load("config.yml")
	if err != nil{
		return nil, err
	}

	serv, err := server.New(cfg)
	if err != nil{
		return nil, err
	}

	return &App{
		loger: loger.New(),
		rpcServer: serv,
		cfg: cfg,
	},nil
}

func (a *App) Run() error {
	a.loger.Info().Msg("starting rpc server...")

	err := a.rpcServer.Run()
	if err != nil{
		return fmt.Errorf("error started server: %v", err)
	}

	a.loger.Info().Msg("Server started succesfully!")
	return nil
}

