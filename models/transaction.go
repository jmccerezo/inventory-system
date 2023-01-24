package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Action int

func (Action) InventoryIn() Action {
	return 1
}
func (Action) InventoryOut() Action {
	return 2
}

type Transaction struct {
	TrackingNumber string
	Item           Item
	ItemID         uint
	Category       Category
	CategoryID     uint
	Quantity       int
	Action         Action
	From           Room
	FromID         uint
	To             Room
	ToID           uint
	Remarks        string
	Date           string
	CreatedBy      string
	CreatedAt      time.Time
	UpdatedBy      string
}

func (t *Transaction) String() string {
	return t.TrackingNumber
}

func (t *Transaction) Save() {
	db := GormDB()
	var count int64
	sum := Summary{}
	db.Preload("Category").Preload("Room").Find(&sum)

	t.TrackingNumber = GenerateTrackingNumber()

	t.Date = time.Now().Format("2006-01-02 - 03:04 PM")

	if t.Action == t.Action.InventoryIn() {
		db.Model(&sum).Where("name = ?", t.Item.Name+"-"+t.To.Name).Count(&count)

		if count > 0 {
			db.Find(&sum, "name=?", t.Item.Name+"-"+t.To.Name).Scan(&sum)
			sum.ItemID = t.ItemID
			sum.CategoryID = t.Item.CategoryID
			sum.Quantity += t.Quantity
			sum.QuantifierID = t.Item.QuantifierID
			sum.LocationID = t.ToID
			db.Save(&sum)

		} else {
			sum.Name = t.Item.Name + "-" + t.To.Name
			sum.ItemID = t.ItemID
			sum.CategoryID = t.Item.CategoryID
			sum.Quantity = t.Quantity
			sum.QuantifierID = t.Item.QuantifierID
			sum.LocationID = t.ToID
			db.Save(&sum)
		}

		t.CategoryID = sum.CategoryID
		db.Save(t)
	}

	if t.Action == t.Action.InventoryOut() {
		db.Model(&sum).Where("name = ?", t.Item.Name+"-"+t.From.Name).Count(&count)

		if count > 0 {
			db.Model(&sum).Where("name = ?", t.Item.Name+"-"+t.To.Name).Count(&count)

			if count > 0 {
				db.Find(&sum, "name=?", t.Item.Name+"-"+t.From.Name).Scan(&sum)
				db.Find(&sum, "name=?", t.Item.Name+"-"+t.From.Name).Scan(&sum)
				sum.Quantity = sum.Quantity - t.Quantity
				sum.CategoryID = t.Item.CategoryID
				db.Save(&sum)

				sum = Summary{}
				db.Find(&sum, "name=?", t.Item.Name+"-"+t.To.Name).Scan(&sum)
				db.Find(&sum, "name=?", t.Item.Name+"-"+t.To.Name).Scan(&sum)
				sum.Quantity = sum.Quantity + t.Quantity
				sum.CategoryID = t.Item.CategoryID
				db.Save(&sum)

			} else {
				db.Find(&sum, "name=?", t.Item.Name+"-"+t.From.Name).Scan(&sum)
				sum.Quantity = sum.Quantity - t.Quantity
				sum.CategoryID = t.Item.CategoryID
				db.Save(&sum)

				sum = Summary{}
				db.Find(&sum, "name=?", t.Item.Name+"-"+t.To.Name).Scan(&sum)
				sum.Name = t.Item.Name + "-" + t.To.Name
				sum.ItemID = t.ItemID
				sum.CategoryID = t.Item.CategoryID
				sum.Quantity = t.Quantity
				sum.QuantifierID = t.Item.QuantifierID
				sum.LocationID = t.ToID
				db.Save(&sum)
			}

			t.CategoryID = sum.CategoryID
			db.Save(t)
		}
	}
}

func GenerateTrackingNumber() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format("20060102150405"))
}

func GormDB() *gorm.DB {
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/inventory_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}
