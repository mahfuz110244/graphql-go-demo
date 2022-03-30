package model

type AuthorGorm struct {
	AuthorID  string `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}

type BookGorm struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Price  int     `json:"price"`
	IsbnNo string  `json:"isbnNo"`
	Author *Author `json:"Author";gorm:"ForeignKey:AuthorID"`
}
