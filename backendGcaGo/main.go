package main

import (
	"backendGcaGo/controllers"
	"backendGcaGo/driver"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()
	//products routes
	router.HandleFunc("/products", controller.GetProducts(db)).Methods("GET")
	router.HandleFunc("/products/{id}", controller.GetProduct(db)).Methods("GET")
	router.HandleFunc("/products", controller.AddProducts(db)).Methods("POST")
	// router.HandleFunc("/products", controller.GetProducts(db)).Methods("PUT")

	//next routes

	fmt.Println("Server running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
