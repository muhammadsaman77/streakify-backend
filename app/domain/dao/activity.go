package dao

import "time"

type Activity struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Target int `json:"target"`
	StartTime time.Time `json:"start_time"`
	RecurrencePeriod int `json:"recurrence_period"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}