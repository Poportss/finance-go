package http

import (
	"github.com/Poportss/finance-go/adapter/http/actuator"
	"github.com/Poportss/finance-go/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Init start the serve
func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransactions)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
