package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jmccerezo/inventory-system/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":8080"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("inventory_system")
	AutoMigrateDB()
	CreateDefaultUser()
	InitialData()
	Handlers()

	http.ListenAndServe(Port, nil)
}

func Handlers() {
	fmt.Println("Handlers")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func GormDB() *gorm.DB {
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/inventory_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func CreateDB(name string) {
	fmt.Println("Database Created")

	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func AutoMigrateDB() {
	fmt.Println("Database Auto Migrated")

	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/inventory_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.Category{},
		&models.Item{},
		&models.Quantifier{},
		&models.Room{},
		&models.Summary{},
		&models.Transaction{},
		&models.User{},
	)
}

func CreateDefaultUser() {
	fmt.Println("Default User Created")
	fmt.Println("Username: admin", "Password: admin")

	db := GormDB()
	user := []models.User{}
	db.Find(&user)

	defaultUser := []models.User{
		{
			Username:  "admin",
			Password:  HashPassword("admin"),
			FirstName: "System",
			LastName:  "Administrator",
			Active:    true,
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Username == users.Username {
				isExisting = true
				break
			}
		}

		if !isExisting {
			db.Create(&defaultUser[i])
		}
	}
}

func HashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func InitialData() {
	fmt.Println("Initial Data Created")
	db := GormDB()

	//Categories
	categories := []models.Category{}
	db.Find(&categories)

	defaultCategory := []models.Category{
		{
			Name: "CATEGORY",
		},
	}

	categoryExist := false
	for i := range defaultCategory {
		categoryExist = false

		for _, category := range categories {
			if defaultCategory[i].Name == category.Name {
				categoryExist = true
				break
			}
		}

		if !categoryExist {
			db.Create(&defaultCategory[i])
		}
	}

	//Rooms
	rooms := []models.Room{}
	db.Find(&rooms)

	defaultRoom := []models.Room{
		{
			Name: "STOCK ROOM",
		},
	}

	isExisting := false
	for i := range defaultRoom {
		isExisting = false

		for _, room := range rooms {
			if defaultRoom[i].Name == room.Name {
				isExisting = true
				break
			}
		}

		if !isExisting {
			db.Create(&defaultRoom[i])
		}
	}

	//Quantifiers
	quantifiers := []models.Quantifier{}
	db.Find(&quantifiers)

	defaultQuantifier := []models.Quantifier{
		{
			Code: "PC",
			Name: "Piece",
		},
		{
			Code: "PAIR",
			Name: "Pair",
		},
		{
			Code: "PCK",
			Name: "Pack",
		},
		{
			Code: "BAG",
			Name: "Bag",
		},
		{
			Code: "BOX",
			Name: "Box",
		},
		{
			Code: "INCH",
			Name: "Inches",
		},
		{
			Code: "FT",
			Name: "Feet",
		},
		{
			Code: "MM",
			Name: "Millimeters",
		},
		{
			Code: "M",
			Name: "Meter",
		},
		{
			Code: "BOTTLE",
			Name: "Bottle",
		},
		{
			Code: "JAR",
			Name: "JAR",
		},
		{
			Code: "CAN",
			Name: "Can",
		},
		{
			Code: "CASE",
			Name: "Case",
		},
		{
			Code: "ROLL",
			Name: "Roll",
		},
		{
			Code: "BNDL",
			Name: "Bundle",
		},
		{
			Code: "L",
			Name: "Liter",
		},
		{
			Code: "G",
			Name: "Gram",
		},
		{
			Code: "KG",
			Name: "Kilogram",
		},
		{
			Code: "UNIT",
			Name: "Unit",
		},
		{
			Code: "SET",
			Name: "Set",
		},
	}

	quantifierExist := false
	for i := range defaultQuantifier {
		quantifierExist = false
		for _, quantifier := range quantifiers {
			if defaultQuantifier[i].Name == quantifier.Name {
				quantifierExist = true
				break
			}
		}

		if !quantifierExist {
			db.Create(&defaultQuantifier[i])
		}
	}
}
