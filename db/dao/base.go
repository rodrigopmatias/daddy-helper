package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type _BaseController struct {
	db *gorm.DB
}

func (c *_BaseController) open() (*gorm.DB, error) {
	if c.db == nil {
		db, err := gorm.Open(sqlite.Open("data.db?_mutex=full&cache=shared&_sync=3&_journal=wal"), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		c.db = db
	}

	return c.db, nil
}
