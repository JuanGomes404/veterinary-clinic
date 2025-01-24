package model

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	name    string `json:"name" gorm:"size:100;not null`
	age     int    `json:"age" gorm: "not null"`
	species string `json:"species" gorm: "size:50";not null`
	owner   string `json:"owner" gorm:"size"`
}
