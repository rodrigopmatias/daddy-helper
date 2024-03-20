package helpers

import (
	"github.com/matias-inc/goenv"
)

type _Config struct {
	TerminalId              string
	MetricAPI               string
	DispatchIntervalSeconds int64
	CollectIntervalSeconds  int64
	DispatchChunkSize       int64
	BusSize                 int64
}

var config *_Config = nil

func GetConfig() *_Config {
	if config == nil {
		config = &_Config{
			TerminalId:              goenv.Config("TERMINAL_ID", "69909214-96e6-fake-b1d0-b5b8029a0faf", goenv.CastString),
			MetricAPI:               goenv.Config("METRIC_API", "http://localhost:3000/v1", goenv.CastString),
			DispatchIntervalSeconds: goenv.Config("DISPATCH_INTERVAL_SECONDS", "300", goenv.CastInt64),
			CollectIntervalSeconds:  goenv.Config("COLLECT_INTERVAL_SECONDS", "60", goenv.CastInt64),
			DispatchChunkSize:       goenv.Config("DISPATCH_CHUNK_SIZE", "20", goenv.CastInt64),
			BusSize:                 goenv.Config("BUS_SIZE", "5", goenv.CastInt64),
		}
	}

	return config
}
