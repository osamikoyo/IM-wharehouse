package data

import (
	"github.com/osamikoyo/IM-wharehouse/internal/data/models"
	"github.com/osamikoyo/IM-wharehouse/internal/updater"
	"github.com/osamikoyo/IM-wharehouse/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
	updater *updater.Updater
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

	go s.updater.Do(product.Count, product.FullCount)

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