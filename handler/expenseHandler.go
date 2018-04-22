package handler

import (
	"github.com/codehell/goAPIExpenses/models"
	"log"
	"encoding/json"
	"net/http"
	"database/sql"
	"fmt"
	"strconv"
)

func ExpenseHandler(db *sql.DB, w http.ResponseWriter, r *http.Request, parts []string) {

	method := r.Method

	expense := models.NewExpense(db)
	if method == "GET" {

		if len(parts) == 3 {
			id, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err = expense.Get(id); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			jsonExpense, err := json.Marshal(expense)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(jsonExpense))
		} else {
			expenses, err := expense.All()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			jsonExpenses, err := json.Marshal(expenses)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(jsonExpenses))
		}

	} else if method == "POST" {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&expense); err != nil {
			log.Fatal(err)
		}
		if err := expense.Create(); err != nil {
			log.Fatal(err)
		}
		jsonExpense, err := json.Marshal(expense)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonExpense)
	}

}