package requests

import (
	"encoding/json"
	"io/ioutil"
)

func save_rating_of_book(b Books) {
	file, _ := json.MarshalIndent(b, "", " ")

	_ = ioutil.WriteFile("books.json", file, 0644)
}

func save_user(u Users) {
	file, _ := json.MarshalIndent(u, "", " ")

	_ = ioutil.WriteFile("users.json", file, 0644)
}
