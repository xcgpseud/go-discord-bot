package entities

import (
	"github.com/jinzhu/gorm"
)

type Pester struct {
	gorm.Model
	UserIdFrom string
	UserIdTo   string
	Message    string
}
