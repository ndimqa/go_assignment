package structs

type Book struct {
	ID             int     `json:"id" gorm:"primaryKey"`
	Title          string  `json:"title"`
	Author         string  `json:"author"`
	Price          float64 `json:"price"`
	Rating         float64 `json:"rating"`
	AmountOfRating int     `json:"amount_of_rating"`
	SumOfRatings   int     `json:"sum_of_ratings"`
	// Comments       []Comment `json:"comments" gorm:"foreignKey:bookId"`
}
