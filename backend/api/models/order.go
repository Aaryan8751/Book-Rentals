package models

import "time"

//Order struct
type Order struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Books       []Books   `json:"books"`
	TotalAmount float64   `json:"total_amount"`
	Time        time.Time `json:"time"`
}
