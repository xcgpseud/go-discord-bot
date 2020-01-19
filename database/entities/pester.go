package entities

import (
	"github.com/jinzhu/gorm"
)

type Pester struct {
	gorm.Model
	UserIdFrom int64
	UserIdTo   int64
	Message    string
}
