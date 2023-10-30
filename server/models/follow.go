package models

import "time"

type Follow struct {
	FollowerId string    `json:"followers_id" gorm:"follower_id"`
	UserId     string    `json:"user_id" gorm:"user_id"`
	CreateAt   time.Time `json:"create_at" gorm:"create_at"`
}

type CreateFollow struct {
	FollowerId string    `json:"followers_id" gorm:"follower_id" binding:"required"`
	UserId     string    `json:"user_id" gorm:"user_id" binding:"required"`
}

func (f *Follow)TableName() string{
	return "follows"
}

func (f *CreateFollow)TableName() string{
	return "follows"
}