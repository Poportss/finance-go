package http

import (
	"github.com/Poportss/finance-go/adapter/http/actuator"
	"github.com/Poportss/finance-go/adapter/http/transaction"
	"net/http"
)

// Init start the serve
func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransactions)

	http.HandleFunc("/health", actuator.Health)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
