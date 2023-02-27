package structs

type Comment struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	UserId int    `json:"user_id"`
	BookId int    `json:"book_id"`
	Text   string `json:"text"`
}
