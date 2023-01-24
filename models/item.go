package models

import (
	"strings"
)

type Item struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Quantifier   Quantifier
	QuantifierID uint
	Category     Category
	CategoryID   uint
}

func (i *Item) String() string {
	return i.Name
}

func (i *Item) Save() {
	db := GormDB()
	i.Name = strings.ToUpper(i.Name)
	db.Save(i)
}
