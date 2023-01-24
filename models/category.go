package models

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func (c *Category) String() string {
	return c.Name
}

func (c *Category) Save() {
	db := GormDB()
	db.Save(c)
}
