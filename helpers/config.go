package helpers

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type _Config struct {
	TerminalId              string `env:"TERMINAL_ID, default=69909214-96e6-fake-b1d0-b5b8029a0faf"`
	MetricAPI               string `env:"METRIC_API, default=http://localhost:3000/v1"`
	DispatchIntervalSeconds int64  `env:"DISPATCH_INTERVAL_SECONDS, default=300"`
	CollectIntervalSeconds  int64  `env:"COLLECT_INTERVAL_SECONDS, default=60"`
	DispatchChunkSize       int64  `env:"DISPATCH_CHUNK_SIZE, default=20"`
	BusSize                 int64  `env:"BUS_SIZE, default=5"`
}

var config *_Config = nil

func GetConfig() *_Config {
	if config == nil {
		var c _Config
		ctx := context.Background()

		if err := envconfig.Process(ctx, &c); err != nil {
			panic(err)
		}

		config = &c
	}

	return config
}
