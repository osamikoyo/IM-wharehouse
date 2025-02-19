package models

type Product struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string
	Count uint
	FullCount uint
	Price uint
}