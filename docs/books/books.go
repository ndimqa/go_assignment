package books

import "math"

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	ID             int     `json:"id" gorm:"primaryKey"`
	Title          string  `json:"title"`
	Author         string  `json:"author"`
	Price          float64 `json:"price"`
	Rating         float64 `json:"rating"`
	AmountOfRating int     `json:"amount_of_rating"`
	SumOfRatings   int     `json:"sum_of_ratings"`
}

type BookInterface interface {
	set_rating(rating int)
}

func (b *Book) set_rating(rating int) {
	b.SumOfRatings = b.SumOfRatings + rating
	b.Rating = math.Floor(float64(b.SumOfRatings) / float64(b.AmountOfRating))
}
