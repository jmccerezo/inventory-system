package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmccerezo/inventory-system/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	data := map[string]interface{}{}

	activeSession := GetActiveSession(r)

	if activeSession != nil {
		http.Redirect(w, r, "/inventory/dashboard", http.StatusSeeOther)
	}

	if r.Method == "POST" {
		dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/pudding_v2?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("Faied to Connect to the Database ", err)
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		user := models.User{}

		db.Where("username = ?", username).Find(&user)

		if CheckPasswordHash(password, user.Password) {
			newSession := uuid.NewString()

			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: newSession,
			})

			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "username",
				Value: user.Username,
			})

			http.Redirect(w, r, "/inventory/dashboard", http.StatusSeeOther)
		}
	}

	data["Title"] = "Inventory System | Login"
	tmpl.Execute(w, data)
}

func GetActiveSession(r *http.Request) *http.Cookie {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key
	}
	return nil
}

func HashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
