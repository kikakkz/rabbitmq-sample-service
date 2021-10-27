package main

import (
	"time"

	cli "github.com/urfave/cli/v2"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
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

func runRabbitMQSample() error {
	err := msgsrv.DeclareQueue("common")
	if err != nil {
		return err
	}

	for {
		logger.Sugar().Infof("Publish hello world")
		err := msgsrv.PublishToQueue("common", "hello world")
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 3)
	}
}
