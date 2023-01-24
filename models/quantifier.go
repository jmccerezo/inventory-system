package models

type Quantifier struct {
	ID   uint `gorm:"primaryKey"`
	Code string
	Name string
}

func (q *Quantifier) String() string {
	return q.Code
}

func (q *Quantifier) Save() {
	db := GormDB()
	db.Save(q)
}
