package models

type Summary struct {
	Name         string
	Item         Item
	ItemID       uint
	Quantifier   Quantifier
	QuantifierID uint
	Category     Category
	CategoryID   uint
	Location     Room
	LocationID   uint
	Quantity     int
}

func (s *Summary) String() string {
	return s.Name

}

func (s *Summary) Save() {
	db := GormDB()
	db.Save(s)
}
