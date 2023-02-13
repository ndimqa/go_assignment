package requests

import (
	"api/docs/books"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) SetData(w http.ResponseWriter, r *http.Request) {
	h.DB.Create(&books.Book{
		ID: 1, Title: "War and Peace", Author: " Leo Tolstoy",
		Price: 12.99, Rating: 12, AmountOfRating: 6, SumOfRatings: 12})
	fmt.Println("COOL")

}

//	func find_book_by_id(id int) (books_pkg.Book, int) {
//		data_books := generate_books().Books
//		for i := 0; i < len(data_books); i++ {
//			n := data_books[i]
//			if n.ID == id {
//				return n, i
//			}
//		}
//		panic("There is no any book with this id")
//	}
//
//	func find_user_by_name(name string) users_pkg.User {
//		users := generate_users().Users
//		for _, n := range users {
//			if n.Username == name {
//				return n
//			}
//		}
//		return User{}
//	}
func (h handler) Get_all_items(w http.ResponseWriter, r *http.Request) {
	var data_books []books.Book
	if result := h.DB.Find(&data_books); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data_books)
}

//
//func Sort_by_price(sortBy string) []Book {
//	data_books := generate_books().Books
//	data := []Book{}
//	data = append([]Book{}, data_books...)
//	if sortBy == "asc" || sortBy == "" {
//		sort.Slice(data, func(i, j int) bool {
//			return data[i].Price < data[j].Price
//		})
//		return data
//	} else if sortBy == "desc" {
//		sort.Slice(data, func(i, j int) bool {
//			return data[i].Price > data[j].Price
//		})
//		return data
//	}
//	fmt.Println("Sort_by_price something went wrong")
//	return data
//}
//
//func Sort_by_rating(sortBy string) []Book {
//	data_books := generate_books().Books
//	data := []Book{}
//	data = append([]Book{}, data_books...)
//	if sortBy == "asc" || sortBy == "" {
//		sort.Slice(data, func(i, j int) bool {
//			return data[i].Price < data[j].Price
//		})
//		return data
//	} else if sortBy == "desc" {
//		sort.Slice(data, func(i, j int) bool {
//			return data[i].Rating > data[j].Rating
//		})
//		return data
//	}
//	fmt.Println("Sort_by_rating something went wrong")
//	return data
//}
//
//func (h handler) Filter(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Filter request sent")
//	sortBy := r.URL.Query().Get("sortBy")           // asc or desc
//	byAttribute := r.URL.Query().Get("byAttribute") // price or rating
//	result := []Book{}
//	if byAttribute == "price" || byAttribute == "" {
//		result = Sort_by_price(sortBy)
//	} else if byAttribute == "rating" {
//		result = Sort_by_rating(sortBy)
//	}
//	json.NewEncoder(w).Encode(result)
//	fmt.Println("Filter request ended")
//}
//
//func (h handler) GiveRating(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("GiveRating request sended")
//	data_books := generate_books().Books
//	x := map[string]int{}
//	b, _ := io.ReadAll(r.Body)
//	err := json.Unmarshal([]byte(b), &x)
//	if err != nil {
//		return
//	}
//	fmt.Println(x["id"])
//	id := x["id"]         // id of book
//	rating := x["rating"] // gaven rating
//	exact_book, book_number := find_book_by_id(id)
//	fmt.Println(exact_book)
//	exact_book.set_rating(rating)
//	fmt.Println(exact_book)
//	data_books[book_number] = exact_book
//	books := Books{data_books}
//	save_rating_of_book(books)
//}
//
//func (h handler) Register(w http.ResponseWriter, r *http.Request) {
//	users := generate_users()
//	fmt.Println("Register user request sended")
//	x := map[string]string{}
//	b, _ := io.ReadAll(r.Body)
//	err := json.Unmarshal([]byte(b), &x)
//	if err != nil {
//		return
//	}
//	if find_user_by_name(x["username"]).Username == "" {
//		fmt.Println("I GOT U BBY")
//		id := len(users.Users)
//		new_user := User{
//			ID:       id,
//			PASSWORD: x["password"],
//			Username: x["username"],
//			Mail:     x["Mail"],
//		}
//		users.Users = append(users.Users, new_user)
//		save_user(users)
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		resp := make(map[string]string)
//		resp["message"] = "Status OK"
//		jsonResp, err := json.Marshal(resp)
//		if err != nil {
//			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
//		}
//		write, err := w.Write(jsonResp)
//		if err != nil {
//			return
//		}
//		return
//	}
//	w.WriteHeader(http.StatusForbidden)
//	w.Header().Set("Content-Type", "application/json")
//	resp := make(map[string]string)
//	resp["message"] = "Change username, it already exists"
//	jsonResp, err := json.Marshal(resp)
//	if err != nil {
//		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
//	}
//	write, err := w.Write(jsonResp)
//	if err != nil {
//		return
//	}
//}
//
//func (h handler) LogIn(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Login user request sended")
//	x := map[string]string{}
//	b, _ := io.ReadAll(r.Body)
//	err := json.Unmarshal([]byte(b), &x)
//	if err != nil {
//		return
//	}
//	username := x["username"]
//	user := find_user_by_name(username)
//	if user.PASSWORD == x["password"] {
//		w.WriteHeader(http.StatusOK)
//		w.Header().Set("Content-Type", "application/json")
//		resp := make(map[string]string)
//		resp["message"] = "Status OK"
//		jsonResp, err := json.Marshal(resp)
//		if err != nil {
//			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
//		}
//		w.Write(jsonResp)
//		return
//	}
//	w.WriteHeader(http.StatusForbidden)
//	w.Header().Set("Content-Type", "application/json")
//	resp := make(map[string]string)
//	resp["message"] = "Status Not OK"
//	jsonResp, err := json.Marshal(resp)
//	if err != nil {
//		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
//	}
//	write, err := w.Write(jsonResp)
//	if err != nil {
//		return
//	}
//}
//
//func (h handler) Find(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Find user request sended")
//	var answer []Book
//	search := strings.ToLower(r.URL.Query().Get("Srch"))
//	data_books := generate_books().Books
//	for i := 0; i < len(data_books); i++ {
//		lower_name := strings.ToLower(data_books[i].Title)
//		if strings.Contains(lower_name, search) == true {
//			answer = append(answer, data_books[i])
//		}
//	}
//	if len(answer) > 0 {
//		json.NewEncoder(w).Encode(answer)
//	}
//	w.WriteHeader(http.StatusForbidden)
//	w.Header().Set("Content-Type", "application/json")
//	resp := make(map[string]string)
//	resp["message"] = "There is nothing"
//	jsonResp, err := json.Marshal(resp)
//	if err != nil {
//		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
//	}
//	write, err := w.Write(jsonResp)
//	if err != nil {
//		return
//	}
//}
