package models

import "time"

type Contract struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	HolderID   uint      `json:"holder_id"`
	PropertyID uint      `json:"property_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	Comments   []Comment `json:"comments"`
}

type Comment struct {
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}
