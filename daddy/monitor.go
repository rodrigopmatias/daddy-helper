package daddy

import (
	"time"

	"github.com/google/uuid"
	"github.com/rodrigopmatias/daddy-helper/db/dao"
	"github.com/rodrigopmatias/daddy-helper/db/input"
)

func writeData(bus <-chan input.Message) {
	for message := range bus {
		dao.MessageController.Create(message)
	}
}

func Monitor() {
	bus := make(chan input.Message, 1)

	go writeData(bus)

	for {
		logger.Info("read enviroment now")
		bus <- input.Message{
			Id:        uuid.NewString(),
			CreatedAt: time.Now().UTC(),
		}
		logger.Info("sleeping")
		time.Sleep(time.Duration(config.CollectIntervalSeconds * int64(time.Second)))
	}
}
