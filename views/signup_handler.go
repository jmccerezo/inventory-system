package views

import (
	"html/template"
	"net/http"

	"github.com/jmccerezo/inventory-system/models"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/signup.html"))
	data := map[string]interface{}{}

	db := models.GormDB()
	user := models.User{}
	firstname := r.FormValue("first-name")
	lastname := r.FormValue("last-name")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if r.Method == "POST" {
		user.FirstName = firstname
		user.LastName = lastname
		user.Username = username
		user.Password = HashPassword(password)
		db.Create(&user)
	}

	data["Title"] = "Inventory System | Signup"
	tmpl.Execute(w, data)
}
