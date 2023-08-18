package personal

import "time"

type Personal struct {
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Telefone  string    `json:"telefone"`
	CreatedAt time.Time `json:"created_at"`
}

type Personals []Personal
