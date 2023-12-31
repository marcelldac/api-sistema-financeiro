package http

import (
	"fmt"
	"net/http"

	"github.com/marcelldac/api-sistema-financeiro/adapter/http/actuator"
	"github.com/marcelldac/api-sistema-financeiro/adapter/http/personal"
	"github.com/marcelldac/api-sistema-financeiro/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func mainResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("===========================API de Sistema financeiro=============================\nOlá, seja bem vindo! Você pode ver os endpoints na documentação. Atualmente o principal é o GET /transactions.\n=================================================================================")
}

func Init() {
	http.HandleFunc("/", mainResponse)
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.PostTransaction)
	http.HandleFunc("/personal", personal.GetPersonal)
	http.HandleFunc("/health", actuator.Health)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
