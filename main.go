package main

import (
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8080", nil)

}

type Transaction struct {
	Title     string
	Amount    float32
	Type      int
	CreatedAt time.Time
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
