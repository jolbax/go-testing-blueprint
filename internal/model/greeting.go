package model

type Greeting struct {
	Language string `json:"language"  binding:"required"`
	Text     string `json:"text"  binding:"required"`
}
