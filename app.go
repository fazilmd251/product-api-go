package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"product-api/models"
)

func startServer(database *sql.DB) error {
	http.HandleFunc("/products", getProducts(database))
	http.HandleFunc("/product/{id}", getProduct(database))
	fmt.Println("Server starting on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func getProducts(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := database.Query("select * from products;")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
			return
		}
		defer rows.Close()

		var products []models.Product

		for rows.Next() {
			var product models.Product
			if err := rows.Scan(
				&product.ProductID,
				&product.ProductName,
				&product.Price,
				&product.StockQuantity,
				&product.CategoryID,
				&product.CreatedAt,
			); err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
				return
			}
			products = append(products, product)
		}

		if err := rows.Err(); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}
func getProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var product models.Product

		err := db.QueryRow("select * from products where product_id=$1", id).Scan(
			&product.ProductID,
			&product.ProductName,
			&product.Price,
			&product.StockQuantity,
			&product.CategoryID,
			&product.CreatedAt,
		)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]error{"error": err})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}
