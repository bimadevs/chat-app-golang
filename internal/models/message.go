package models

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey"`
	Sender    string    `json:"sender"`
	Message   string    `json:"message"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
