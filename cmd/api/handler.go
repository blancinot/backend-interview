package main

import (
	"net/http"

	"github.com/gustvision/backend-interview/pkg/account"
	"github.com/gustvision/backend-interview/pkg/user"
)

type handler struct {
	user    user.App
	account account.App
}

func (h handler) listen(host string) error {
	http.HandleFunc("/user", h.GetUser)
	http.HandleFunc("/transaction", h.CreateTransaction)
	http.HandleFunc("/max-transaction", h.GetMaxTransaction)

	if err := http.ListenAndServe(host, nil); err != nil {
		return err
	}

	return nil
}
