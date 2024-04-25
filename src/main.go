package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"context"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func getProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	cacheResult, err := rdb.Get(ctx, "product:"+vars["id"]).Result()
	if err == redis.Nil {
		log.Println("Product not found in cache")
	} else if err != nil {
		http.Error(w, "Error fetching product from cache", http.StatusInternalServerError)
		return
	} else {
		log.Println("Product found in cache")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cacheResult))
		return
	}

	db, err := sql.Open("mysql", "root:qwertyios02@tcp(127.0.0.1:3306)/assigment_2")
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var product Product
	err = db.QueryRow("SELECT id, name, description, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		http.Error(w, "Product not found in database", http.StatusNotFound)
		return
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Error converting product data to JSON", http.StatusInternalServerError)
		return
	}

	err = rdb.Set(ctx, "product:"+vars["id"], productJSON, 0).Err()
	if err != nil {
		http.Error(w, "Error caching product data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productJSON)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", getProductByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
