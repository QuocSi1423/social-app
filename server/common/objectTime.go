package common

import "time"

type ObjectTime struct {
	CreateAt time.Time `json:"create_at" gorm:"create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
}