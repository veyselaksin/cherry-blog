
package models

import (
	"time"
)

type Blog struct {
	ID            int       `gorm:"primaryKey;autoIncrement:true;not null:true;index;" json:"id"`
	CreatedDate   time.Time `gorm:"default:current_timestamp" json:"created_date"`
	UpdatedDate   time.Time `gorm:"default:current_timestamp" json:"updated_date"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
}

func (c Blog) TableName() string {
	return "public.blog"
}
