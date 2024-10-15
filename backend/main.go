package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var store = sessions.NewCookieStore([]byte("super-secret-key")) // Session store

type Godown struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_godown"`
}

type Item struct {
	ItemID     string  `json:"item_id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Category   string  `json:"category"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"`
	GodownID   string  `json:"godown_id"`
	Brand      string  `json:"brand"`
	Attributes string  `json:"attributes"`
	ImageURL   string  `json:"image_url"`
}

func getGodowns(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, parent_godown FROM godowns")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var godowns []Godown
	for rows.Next() {
		var godown Godown
		if err := rows.Scan(&godown.ID, &godown.Name, &godown.ParentID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		godowns = append(godowns, godown)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(godowns)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ItemID, &item.Name, &item.Quantity, &item.Category,
			&item.Price, &item.Status, &item.GodownID, &item.Brand, &item.Attributes, &item.ImageURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItemsByGodown(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	godownID := params["godown_id"]

	rows, err := db.Query("SELECT * FROM items WHERE godown_id = ?", godownID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ItemID, &item.Name, &item.Quantity, &item.Category,
			&item.Price, &item.Status, &item.GodownID, &item.Brand, &item.Attributes, &item.ImageURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	initDB()

	router := mux.NewRouter()

	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")


	// api := router.PathPrefix("/api").Subrouter()
	// api.Use(authMiddleware)
	router.HandleFunc("/api/godowns", getGodowns).Methods("GET")
	router.HandleFunc("/api/items", getAllItems).Methods("GET")
	router.HandleFunc("/items/{godown_id}", getItemsByGodown).Methods("GET")

	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
