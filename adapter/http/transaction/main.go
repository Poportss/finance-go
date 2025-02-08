package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/Poportss/finance-go/model/transaction"
	"github.com/Poportss/finance-go/util"
	"io"
	"net/http"
)

// GetTransactions find some transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Value:     1200.1,
			Type:      0,
			CreatedAt: util.StringToTime("2025-02-08T15:04:05"),
		},
	}

	err := json.NewEncoder(w).Encode(transactions)
	if err != nil {
		return
	}

}

// CreateATransactions create a new transaction
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
