package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Long  string
	Short string
}
