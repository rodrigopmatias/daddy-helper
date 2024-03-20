package models

type Message struct {
	Id           string `json:"id" gorm:"type:CHAR(36);primaryKey"`
	CreatedAt    int64  `json:"createdAt" gorm:"notNull"`
	DispatchedAt int64  `json:"dispatchedAt"`
}
