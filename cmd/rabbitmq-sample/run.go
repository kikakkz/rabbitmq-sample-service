package main

import (
	"time"

	cli "github.com/urfave/cli/v2"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	msgcli "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/client"
	"github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/common" //nolint
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		logger.Sugar().Infof("Run rabbitmq sample")
		return runRabbitMQSample()
	},
}

func consumer(id int) {
	_cli, err := msgcli.New(rabbitmq.MyServiceNameToVHost())
	if err != nil {
		logger.Sugar().Errorf("fail to create rabbitmq client: %v", err)
		return
	}

	msgs, err := _cli.Consume("common")
	if err != nil {
		logger.Sugar().Errorf("fail to consume msgs: %v", err)
		return
	}

	for d := range msgs {
		logger.Sugar().Infof("consume '%s' by %v", d.Body, id)
	}
}

func runRabbitMQSample() error {
	err := msgsrv.DeclareQueue("common")
	if err != nil {
		return err
	}

	go consumer(1)
	go consumer(2)

	for {
		logger.Sugar().Infof("Publish hello world")
		err := msgsrv.PublishToQueue("common", "hello world")
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 3)
	}
}
