package models

import "time"

type ReplyComment struct {
	Id                string    `json:"id" gorm:"id"`
	PostId            string    `json:"post_id" gorm:"post_id"`
	Content           string    `json:"content" gorm:"content"`
	ImageUrl          string    `json:"image_url" gorm:"image_url"`
	UserId            string    `json:"user_id" gorm:"user_id"`
	CountReplyComment int       `json:"count_reply_comment" gorm:"count_reply_comment"`
	CreateAt          time.Time `json:"create_at" gorm:"create_at"`
	UpdateAt          time.Time `json:"update_at" gorm:"update_at"`
	ParentCommentId   string    `json:"parent_comment_id" gorm:"parent_comment_id"`
}

type ReplyCommentUpdate struct{
	Id                string    `json:"id" gorm:"id"`
	Content           string    `json:"content" gorm:"content"`
	ImageUrl          string    `json:"image_url" gorm:"image_url"`
	UpdateAt          time.Time `json:"update_at" gorm:"update_at"`
}

func (c *ReplyComment)TableName() string{
	return "reply_comment"
}

func (c *ReplyCommentUpdate)TableName() string{
	return "reply_comment"
}