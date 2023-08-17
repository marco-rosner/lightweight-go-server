package models

type Person struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required,len=32"`
	Nickname string `json:"nickname" binding:"required,len=100"`
	Birth    string `json:"birh" binding:"required"`
}
