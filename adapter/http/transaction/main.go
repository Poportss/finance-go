package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/Poportss/finance-go/model/transaction"
	"io"
	"net/http"
	"time"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2025-02-08T15:04:05")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Value:     1200.1,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	err := json.NewEncoder(w).Encode(transactions)
	if err != nil {
		return
	}

}

func CreateATransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions transaction.Transactions

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &transactions)

	fmt.Printf("%+v\n", transactions)

}
