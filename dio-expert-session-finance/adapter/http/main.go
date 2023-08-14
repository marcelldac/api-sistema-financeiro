package http

import (
	"net/http"

	"github.com/marcelldac/praticando-golang/dio-expert-session-finance/adapter/http/actuator"
	"github.com/marcelldac/praticando-golang/dio-expert-session-finance/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)
	http.HandleFunc("/health", actuator.Health)
	http.ListenAndServe(":8080", nil)
}
