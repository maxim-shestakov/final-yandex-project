package models

import (
	"time"
)

type Task struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
	Repeat string `json:"repeat"`
}