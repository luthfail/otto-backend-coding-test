package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Biller struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Category    string    `gorm:"size:255;not null;" json:"category"`
	Product     string    `gorm:"size:255;not null;" json:"product"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	Price       int       `gorm:"not null;" json:"price"`
	Fee         int       `gorm:"not null;" json:"fee"`
	TotalPay    int       `gorm:"not null;" json:"totalPay"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Biller) Prepare() {
	p.ID = 0
	p.Category = html.EscapeString(strings.TrimSpace(p.Category))
	p.Product = html.EscapeString(strings.TrimSpace(p.Product))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Price = 0
	p.Fee = 0
	p.TotalPay = 0
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Biller) Validate() error {

	if p.Category == "" {
		return errors.New("Required Category")
	}
	if p.Product == "" {
		return errors.New("Required Product")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}
	return nil
}

func (p *Biller) SaveBiller(db *gorm.DB) (*Biller, error) {
	var err error
	err = db.Debug().Model(&Biller{}).Create(&p).Error
	if err != nil {
		return &Biller{}, err
	}
	return p, nil
}

func (p *Biller) FindAllBillers(db *gorm.DB) (*[]Biller, error) {
	var err error
	Billers := []Biller{}
	err = db.Debug().Model(&Biller{}).Limit(100).Find(&Billers).Error
	if err != nil {
		return &[]Biller{}, err
	}
	return &Billers, nil
}

func (p *Biller) FindBillerByID(db *gorm.DB, pid uint64) (*Biller, error) {
	var err error
	err = db.Debug().Model(&Biller{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Biller{}, err
	}
	return p, nil
}

func (p *Biller) UpdateABiller(db *gorm.DB) (*Biller, error) {

	var err error

	err = db.Debug().Model(&Biller{}).Where("id = ?", p.ID).Updates(Biller{Category: p.Category, Product: p.Product, Description: p.Description, Price: p.Price, Fee: p.Fee, TotalPay: p.TotalPay, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Biller{}, err
	}
	return p, nil
}

func (p *Biller) DeleteABiller(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Biller{}).Where("id = ? and author_id = ?", pid, uid).Take(&Biller{}).Delete(&Biller{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Biller not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
