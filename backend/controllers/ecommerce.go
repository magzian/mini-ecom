package controllers

/* import (
	"backend/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var products []models.Product
		db.Find(&products)
		json.NewEncoder(w).Encode(products)
	}
}

func CreateOrder(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		json.NewDecoder(r.Body).Decode(&order)
		db.Create(&order)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(order)
	}
} */
