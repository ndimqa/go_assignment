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
	router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "front/index.html")
	})
	router.HandleFunc("/structs", h.GetAllItems).Methods("GET")    // get all items
	router.HandleFunc("/filter", h.Filter).Methods("GET")          // filtering by rating or price
	router.HandleFunc("/register", h.Register).Methods("POST")     // register
	router.HandleFunc("/login", h.LogIn).Methods("GET")            // authorization
	router.HandleFunc("/search", h.Find).Methods("GET")            // search book by name
	router.HandleFunc("/give_rating", h.GiveRating).Methods("PUT") // giving rating
	router.HandleFunc("/set_data", h.SetData).Methods("POST")
	router.HandleFunc("/publish", h.PublishBook).Methods("POST")
	router.HandleFunc("/purchase", h.Purchase).Methods("GET")
	// commenting
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
