package model

import (
	"time"
)

type CreateRequest struct {
	Author   int64
	ChatName string
	Users    []int64
}

type SendMessageRequest struct {
	From      int64
	ChatId    int64
	Text      string
	Timestamp time.Time
}
