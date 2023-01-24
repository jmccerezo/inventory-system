package models

type User struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Active    bool
}

func (u *User) Save() {
	db := GormDB()
	db.Save(u)
}
