package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Type        string    `json:"type"`
	StartTime   time.Time `json:"starttime"`
	EndTime     time.Time `json:"endtime"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
}
