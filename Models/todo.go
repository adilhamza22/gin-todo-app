package Models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
