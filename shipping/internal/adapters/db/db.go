package db_adapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kaua-victor/microservices/shipping/internal/application/core/domain"
)

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.Shipping{})
	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(s *domain.Shipping) error {
	return a.db.Create(s).Error
}
