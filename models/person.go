package models

import "time"

type Person struct {
	ID        string `json:"id" `
	Name      string `json:"name" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Birth     string `json:"birh" binding:"required"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
