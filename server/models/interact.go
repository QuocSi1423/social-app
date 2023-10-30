package models

import "time"

type Interact struct {
	PostId   string    `json:"post_id" gorm:"post_id"`
	UserId   string    `json:"user_id" gorm:"user_id"`
	CreateAt time.Time `json:"create_at" gorm:"create_at"`
}
