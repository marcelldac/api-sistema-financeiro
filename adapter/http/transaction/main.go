package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/marcelldac/api-sistema-financeiro/model/transaction"
	"github.com/marcelldac/api-sistema-financeiro/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-04-05T11:45:26"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = io.ReadAll(r.Body)

	json.Unmarshal(body, &res)

	fmt.Print(res)
}
