package model

import "gorm.io/gorm"

type Pet struct {
	gorm.Model
	Name    string `json:"Name" gorm:"size:100;not null`
	Age     int    `json:"Age" gorm: "not null"`
	Species string `json:"Species" gorm: "size:50";not null`
	Owner   string `json:"Owner" gorm:"size"`
}
