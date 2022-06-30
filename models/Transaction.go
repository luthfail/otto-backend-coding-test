package models

import (
	"time"
)

type Transaction struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserId    uint32 `gorm:"not null" json:"user_id"`
	BillerId  uint32 `gorm:"not null" json:"biller_id"`
	isPay     bool
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
