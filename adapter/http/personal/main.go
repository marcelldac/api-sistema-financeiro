package personal

import (
	"encoding/json"
	"net/http"

	"github.com/marcelldac/praticando-golang/model/personal"
	"github.com/marcelldac/praticando-golang/util"
)

func GetPersonal(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var personals = personal.Personals{
		personal.Personal{
			Nome:      "Marcell",
			Email:     "marcell@gmail.com",
			Telefone:  "+55(71)99999-9999",
			CreatedAt: util.StringToTime("2023-08-18T18:52:26"),
		},
	}

	_ = json.NewEncoder(w).Encode(personals)
}
