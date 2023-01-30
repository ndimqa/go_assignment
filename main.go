package main

import (
	"api/requests"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", requests.Get_all_items).Methods("GET")    // get all items
	router.HandleFunc("/filter", requests.Filter).Methods("GET")          // filtering by rating or price
	router.HandleFunc("/register", requests.Register).Methods("POST")     // register
	router.HandleFunc("/login", requests.LogIn).Methods("GET")            // authorization
	router.HandleFunc("/search", requests.Find).Methods("GET")            // search book by name
	router.HandleFunc("/give_rating", requests.GiveRating).Methods("PUT") // giving rating
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println("API started")
	handleRequests()
}
