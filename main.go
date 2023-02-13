package main

import (
	"api/docs/db"
	"api/requests"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	DB := db.Init()
	h := requests.New(DB)
	fmt.Println("API started")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", h.Get_all_items).Methods("GET") // get all items
	//router.HandleFunc("/filter", requests.Filter).Methods("GET")          // filtering by rating or price
	//router.HandleFunc("/register", requests.Register).Methods("POST")     // register
	//router.HandleFunc("/login", requests.LogIn).Methods("GET")            // authorization
	//router.HandleFunc("/search", requests.Find).Methods("GET")            // search book by name
	//router.HandleFunc("/give_rating", requests.GiveRating).Methods("PUT") // giving rating
	router.HandleFunc("/set_data", h.SetData).Methods("POST")
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
