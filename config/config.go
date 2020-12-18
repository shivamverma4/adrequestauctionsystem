package config

import (
	"github.com/op/go-logging"

	"adrequestauctionsystem/internal/configuration"
)

var config configuration.Config

const namespace = "adRequestAuctionSystem"

func init() {
	config = configuration.Config{
		Namespace:          namespace,
		Deployment:         configuration.DEBUG,
		LogLevel:           logging.INFO,
		LogFilePath:        "../internal/logs/adRequestAuctionSystem.log",
		RequestLogFilePath: "../internal/logs/request.log",
		Port:               ApiPort,
	}
}

func GetConfig() *configuration.Config {
	return &config
}
