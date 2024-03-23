package daddy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rodrigopmatias/daddy-helper/db/dao"
	"github.com/rodrigopmatias/daddy-helper/db/models"
)

type Metric struct {
	TerminalId string `json:"terminalId"`
	CreatedAt  int64  `json:"createdAt"`
}

func readMessages(messages <-chan models.Message) {
	for message := range messages {
		logger.Infof("%s - dispatched", message.Id)
		values := map[string]interface{}{
			"dispatched_at": time.Now().UTC().Unix(),
		}

		err := dispatchMessage(Metric{config.TerminalId, message.CreatedAt})
		if err != nil {
			logger.Info(err)
		} else {
			dao.MessageController.Update(message.Id, values)
		}
	}
}

func dispatchMessage(metric Metric) error {
	body, err := json.Marshal(metric)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", config.MetricAPI, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		return fmt.Errorf("bad status code (%d) in response", res.StatusCode)
	}

	return nil
}

func Dispatch() {
	bus := make(chan models.Message, config.BusSize)

	go readMessages(bus)

	for {
		logger.Info("dispatch acumulated messages ...")

		messages, err := dao.MessageController.ListNotDispatched(0, int(config.DispatchChunkSize))
		if err != nil {
			logger.Info(err)
		}

		for _, message := range messages {
			bus <- message
		}

		time.Sleep(time.Duration(config.DispatchIntervalSeconds * int64(time.Second)))
	}
}
