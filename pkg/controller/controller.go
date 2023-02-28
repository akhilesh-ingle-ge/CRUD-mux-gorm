package controller

import (
	"crud/pkg/config"
	"crud/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)
	config.DB.Create(&order)
	json.NewEncoder(w).Encode(order)
	// fmt.Println("Hello")
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	Vars := mux.Vars(r)
	id := Vars["id"]
	config.DB.Preload("Items").First(&order, id)
	json.NewEncoder(w).Encode(&order)
	// fmt.Println("Hello")
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []models.Order
	config.DB.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(&orders)
	// fmt.Println("Hello")
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	Vars := mux.Vars(r)
	id := Vars["id"]
	config.DB.Preload("Items").First(&order, id)
	// config.DB.First(&order, id)
	json.NewDecoder(r.Body).Decode(&order)
	config.DB.Preload("Items").Save(&order)
	json.NewEncoder(w).Encode(order)
	fmt.Println("Hello")
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Vars := mux.Vars(r)
	// id := Vars["id"]
	id_demo := Vars["id"]
	id, _ := strconv.Atoi(id_demo)
	// config.DB.Preload("Item").Where("order_id = ?", id).Delete(&models.Order{})
	config.DB.Where("order_id = ?", id).Delete(&models.Item{})
	config.DB.Where("order_id = ?", id).Delete(&models.Order{})
	// fmt.Println("Hello")
}
