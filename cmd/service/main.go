package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rodrigopmatias/daddy-helper/daddy"
	"github.com/rodrigopmatias/daddy-helper/helpers"
)

var (
	logger = helpers.GetLogger()
	config = helpers.GetConfig()
)

func printConfig() {
	logger.Info("TerminalID ..........: ", config.TerminalId)
	logger.Info("Metric API ..........: ", config.MetricAPI)
	logger.Info("BUS Size ............: ", config.BusSize)
	logger.Info("Dispatch chunk size  : ", config.DispatchChunkSize)
	logger.Info("Collect .............: ", config.CollectIntervalSeconds, " segundos")
	logger.Info("Dispatch ............: ", config.DispatchIntervalSeconds, " segundos")
}

func main() {
	printConfig()

	done := make(chan bool, 1)
	signalbus := make(chan os.Signal, 1)
	signal.Notify(signalbus, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("service is running now ...")
		logger.Info("waiting for signals ...")
		sig := <-signalbus

		logger.Info("recived signal: ", sig)
		done <- true
	}()

	go daddy.Monitor()
	go daddy.Dispatch()

	<-done
	logger.Info("finish service")
}
