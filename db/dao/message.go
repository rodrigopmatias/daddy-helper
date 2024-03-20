package dao

import (
	"net/http"
	"time"

	"github.com/rodrigopmatias/daddy-helper/db/input"
	"github.com/rodrigopmatias/daddy-helper/db/models"
	"gorm.io/gorm"
)

type _MessageController struct {
	_BaseController
}

func (c _MessageController) Update(id string, data map[string]interface{}) (int, *DAOError) {
	db, err := c.open()
	if err != nil {
		return 0, NewDAOError("cant open database connection", http.StatusInternalServerError)
	}

	message := models.Message{}
	db.Transaction(func(_db *gorm.DB) error {
		tx := _db.Model(&message).Where(models.Message{Id: id}).UpdateColumns(data)
		if tx.Error != nil {
			return tx.Error
		}

		return nil
	})

	return 0, nil
}

func (c _MessageController) ListNotDispatched(offset int, limit int) ([]models.Message, *DAOError) {
	db, err := c.open()
	if err != nil {
		return nil, NewDAOError("cant open database connection", http.StatusInternalServerError)
	}

	entities := make([]models.Message, 0)
	tx := db.Offset(offset).Limit(limit).Order("created_at ASC").Where("dispatched_at = ?", 0).Find(&entities)
	if tx.Error != nil {
		return nil, NewDAOError(tx.Error.Error(), http.StatusInternalServerError)
	}

	return entities, nil
}

func (c _MessageController) Create(data input.Message) (*models.Message, *DAOError) {
	db, err := c.open()
	if err != nil {
		return nil, NewDAOError("cant open database connection", http.StatusInternalServerError)
	}

	entity := &models.Message{
		Id:           data.Id,
		CreatedAt:    time.Now().UTC().Unix(),
		DispatchedAt: 0,
	}

	err = db.Transaction(func(_db *gorm.DB) error {
		tx := db.Create(entity)
		return tx.Error
	})

	if err != nil {
		return nil, NewDAOError(err.Error(), http.StatusUnprocessableEntity)
	}

	return entity, nil
}
