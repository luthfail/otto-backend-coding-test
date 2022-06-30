package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Balance struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	User      User      `json:"User"`
	UserId    uint32    `gorm:"not null" json:"user_id"`
	Wallet    int       `gorm:"not null;" json:"balance"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Balance) Prepare() {
	p.ID = 0
	p.User = User{}
	p.Wallet = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Balance) SaveBalance(db *gorm.DB) (*Balance, error) {
	var err error
	err = db.Debug().Model(&Balance{}).Create(&p).Error
	if err != nil {
		return &Balance{}, err
	}
	return p, nil
}
