package data

import (
	"github.com/osamikoyo/IM-wharehouse/internal/config"
	"github.com/osamikoyo/IM-wharehouse/internal/data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	db *gorm.DB
}

func New(cfg *config.Config) (*Storage, error){
	db, err := gorm.Open(sqlite.Open(cfg.DSN))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) RezervProduct(name string) error {
	var product models.Product

	result := s.db.Where(&models.Product{Name: name}).Find(&product)
	if result.Error != nil{
		return result.Error
	}

	result = s.db.Model(&product).Update("Count", product.Count-1)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Storage) AddProductCount(count uint, name string) error {
	var product models.Product

	result := s.db.Where(&models.Product{Name: name}).Find(&product)
	if result.Error != nil{
		return result.Error
	}

	result = s.db.Model(&product).Update("Count", product.Count+count)
	if result.Error != nil {
		return result.Error
	}

	return nil
}