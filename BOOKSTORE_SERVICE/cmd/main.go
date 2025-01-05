package main

import (
	"log"
	"net/http"

	"github.com/Wandhey/BookStores/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:9010", r))
	err := http.ListenAndServe("localhost:9010", r)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
