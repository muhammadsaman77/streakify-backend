package dao

import "time"

type ActivityHistory struct {
	ID int `json:"id"`
	ActivityId int `json:"activity_id"`
	Achieved int `json:"target"`
	AchievedAt time.Time `json:"achieved_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
}