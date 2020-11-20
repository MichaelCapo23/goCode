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

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()
	//products routes
	router.HandleFunc("/products", controller.GetProducts(db)).Methods("GET")
	router.HandleFunc("/products/{id}", controller.GetProduct(db)).Methods("GET")
	router.HandleFunc("/products", controller.AddProducts(db)).Methods("POST")
	router.HandleFunc("/products", controller.EditProducts(db)).Methods("PUT")

	//discounts routes
	router.HandleFunc("/discounts", controller.GetDiscounts(db)).Methods("GET")
	router.HandleFunc("/discounts/{id}", controller.GetDiscount(db)).Methods("GET")
	router.HandleFunc("/discounts", controller.EditDiscount(db)).Methods("PUT")

	//rewards
	router.HandleFunc("/rewards/{id}", controller.GetReward(db)).Methods("GET")

	//favorties
	router.HandleFunc("/favorites", controller.GetFavorites(db)).Methods("GET")
	router.HandleFunc("/favorites", controller.AddFavorites(db)).Methods("POST")
	router.HandleFunc("/favorites", controller.EditFavorites(db)).Methods("PUT")

	fmt.Println("Server running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
