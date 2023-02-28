package main

import (
	"crud/pkg/config"
	"crud/pkg/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.SetUpDB()

	router := mux.NewRouter()

	//Create
	router.HandleFunc("/orders", controller.CreateOrder).Methods("POST")
	//Read
	router.HandleFunc("/order/{id}", controller.GetOrder).Methods("GET")
	//Read-All
	router.HandleFunc("/orders", controller.GetOrders).Methods("GET")
	//Update
	router.HandleFunc("/orders/{id}", controller.UpdateOrder).Methods("PATCH")
	//Delete
	router.HandleFunc("/orders/{id}", controller.DeleteOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
