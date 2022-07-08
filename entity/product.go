package entity

import "time"

type Product struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Date        time.Time `json:"date"`
}
