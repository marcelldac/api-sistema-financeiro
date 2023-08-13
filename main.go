package main

import (
	"net/http"

	"github.com/marcelldac/praticando-golang/dio-expert-session-finance/adapter/http/transaction"
)

func main() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.ListenAndServe(":8080", nil)

}
