package models

import "time"

type Post struct {
	Id            string    `json:"id" gorm:"id"`
	UserId        string    `json:"user_id" gorm:"user_id"`
	Title         string    `json:"title" gorm:"title"`
	Description   string    `json:"description" gorm:"description"`
	ImageUrl      string    `json:"image_url" gorm:"image_url"`
	CountInteract int       `json:"count_interact" gorm:"count_interact"`
	CountComment  int       `json:"count_comment" gorm:"count_comment"`
	CreateAt      time.Time `json:"create_at" gorm:"create_at"`
}

type PostCreate struct{
	Id            string    `json:"id" gorm:"id"`
	UserId        string    `json:"user_id" gorm:"user_id"`
	Title         string    `json:"title" gorm:"title"`
	Description   string    `json:"description" gorm:"description"`
	ImageUrl      string    `json:"image_url" gorm:"image_url"`
}

func (p *Post)TableName() string{
	return "posts"
}

func (p *PostCreate)TableName() string{
	return "posts"
}