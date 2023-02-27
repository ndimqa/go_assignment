package requests

import (
	"api/docs/structs"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func (h handler) WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST")
	http.ServeFile(w, r, "front/index.html")
}

func (h handler) SetData(w http.ResponseWriter, r *http.Request) {
	h.DB.Create(&structs.Book{
		ID: 1, Title: "War and Peace", Author: " Leo Tolstoy",
		Price: 12.99, Rating: 12, AmountOfRating: 6, SumOfRatings: 12})
	fmt.Println("COOL")
}

func (h handler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	var data_books []structs.Book
	if result := h.DB.Find(&data_books); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data_books)
}

func (h handler) SortByPrice(sortBy string) []structs.Book {
	var data []structs.Book
	if sortBy == "asc" || sortBy == "" {
		h.DB.Order("price asc").Find(&data)
		return data
	} else if sortBy == "desc" {
		h.DB.Order("price desc").Find(&data)
		return data
	}
	fmt.Println("Sort_by_price something went wrong")
	return data
}

func (h handler) SortByRating(sortBy string) []structs.Book {
	var data []structs.Book
	if sortBy == "asc" || sortBy == "" {
		h.DB.Order("rating asc").Find(&data)
		return data
	} else if sortBy == "desc" {
		h.DB.Order("price desc").Find(&data)
		return data
	}
	fmt.Println("Sort_by_rating something went wrong")
	return data
}

func (h handler) Filter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Filter request sent")
	sortBy := r.URL.Query().Get("sortBy")           // asc or desc
	byAttribute := r.URL.Query().Get("byAttribute") // price or rating
	var result []structs.Book
	if byAttribute == "price" || byAttribute == "" {
		result = h.SortByPrice(sortBy)
	} else if byAttribute == "rating" {
		result = h.SortByRating(sortBy)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	fmt.Println("Filter request ended")
}

func (h handler) GiveRating(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GiveRating request sent")
	var book structs.Book
	x := map[string]int{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal([]byte(b), &x)
	if err != nil {
		return
	}
	fmt.Println(x["id"])
	fmt.Println(x["rating"])
	id := x["id"]         // id of book
	rating := x["rating"] // gaven rating
	h.DB.First(&book, "id = ?", id)
	book.AmountOfRating = book.AmountOfRating + 1
	book.SumOfRatings = book.SumOfRatings + rating
	fmt.Println(math.Round(float64(book.SumOfRatings / book.AmountOfRating)))
	book.Rating = math.Round(float64(book.SumOfRatings / book.AmountOfRating))
	h.DB.Save(book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h handler) Find(w http.ResponseWriter, r *http.Request) {
	var books_list []structs.Book
	fmt.Println("Find user request sended")
	search := r.URL.Query().Get("Srch")
	search = "%" + search + "%"
	h.DB.Where("title LIKE ?", search).First(&books_list)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books_list)
}

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register user request sent")
	var user structs.User
	x := map[string]string{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal([]byte(b), &x)
	if err != nil {
		return
	}

	h.DB.First(&user, "username = ?", x["username"])
	if user.Username == "" {
		fmt.Println("I GOT U BBY")
		newUser := structs.User{
			PASSWORD: x["password"],
			Username: x["username"],
			Mail:     x["mail"],
			Role:     0,
		}
		h.DB.Save(&newUser)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Change username, it already exists"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
}

func (h handler) LogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login user request sent")
	x := map[string]string{}
	b, _ := io.ReadAll(r.Body)
	var user structs.User
	err := json.Unmarshal([]byte(b), &x)
	if err != nil {
		return
	}
	h.DB.First(&user, "username = ?", x["username"])
	if user.PASSWORD == x["password"] {
		fmt.Println("Successfully Logged In")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Not OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
}

func (h handler) PublishBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Publish request sent")
	var user structs.User
	x := map[string]string{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal([]byte(b), &x)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status Not OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			return
		}
	}
	h.DB.First(&user, "username = ?", x["username"])
	if user.Role != 0 {
		price, err := strconv.ParseFloat(x["price"], 8)
		book := structs.Book{
			Title:          x["title"],
			Author:         x["author"],
			Price:          price,
			Rating:         0,
			AmountOfRating: 0,
			SumOfRatings:   0,
		}
		h.DB.Save(&book)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "You are not allowed to do this"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
}

func (h handler) Purchase(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Purchase request sent")
	x := map[string][]int{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal([]byte(b), &x)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status Not OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			return
		}
	}
	var books []structs.Book
	var books_titles []string
	var book structs.Book
	for _, value := range x["books_id"] {
		h.DB.First(&book, "id = ?", value)
		books = append(books, book)
		books_titles = append(books_titles, book.Title)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	string_response := strings.Join(books_titles, ", ")
	resp["message"] = string_response
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
}

func (h handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateComment request sent")
	var user structs.User
	var book structs.Book
	x := map[string]string{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal([]byte(b), &x)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status Not OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			return
		}
	}
	book_id, book_err := strconv.Atoi(x["book_id"])

	if book_err != nil {
		fmt.Println(book_err)
		fmt.Println("Error during book conversion")
		return
	}
	user_id, user_err := strconv.Atoi(x["user_id"])

	if user_err != nil {
		fmt.Println("Error during user conversion")
		return
	}
	h.DB.First(&book, "id = ?", book_id)
	h.DB.First(&user, "id = ?", user_id)
	var count int64
	h.DB.Table("comments").Count(&count)
	comment := structs.Comment{
		Id:     int(count),
		UserId: user_id,
		BookId: book_id,
		Text:   x["comment_text"],
	}
	book.Comments = append(book.Comments, comment)
	h.DB.Save(&book)
	h.DB.Create(comment)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		return
	}
}
