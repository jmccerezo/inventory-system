package models

type Room struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (r *Room) String() string {
	return r.Name
}

func (r *Room) Save() {
	db := GormDB()
	db.Save(r)
}
