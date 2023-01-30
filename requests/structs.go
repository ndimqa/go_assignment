package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type Books struct {
	Books []Book `json: "books"`
}

type Users struct {
	Users []User `json: "user"`
}

type Book struct {
	ID            int     `json: "id"`
	Title         string  `json: "title"`
	Author        string  `json: "author"`
	Price         float64 `json: "price"`
	Rating        float64 `json: "rating"`
	AmounOfRating int     `json: "amouunt_of_rating"`
	ListOfRatings []int   `json: "list_of_ratings"`
}

type User struct {
	ID       int    `json: "id"`
	PASSWORD string `json: "password"`
	Username string `json: "username"`
	Mail     string `json: "mail"`
}

type BookInterface interface {
	set_rating(rating int)
}

func (b *Book) set_rating(rating int) {
	total_amount := 0
	b.ListOfRatings = append(b.ListOfRatings, rating)
	for i := 0; i < len(b.ListOfRatings); i++ {
		total_amount += b.ListOfRatings[i]
	}
	b.AmounOfRating += 1
	b.Rating = math.Floor(float64(total_amount) / float64(b.AmounOfRating))
}

func generate_books() Books {
	fileContent, err := os.Open("books.json")

	if err != nil {
		log.Fatal(err)
		panic("Something went wrong with books.json")
	}

	fmt.Println("generated Books")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var data_books Books

	json.Unmarshal(byteResult, &data_books)
	return data_books
}

func generate_users() Users {
	fileContent, err := os.Open("users.json")

	if err != nil {
		log.Fatal(err)
		panic("Something went wrong with books.json")
	}

	fmt.Println("generated Users")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var data_users Users

	json.Unmarshal(byteResult, &data_users)
	return data_users
}
