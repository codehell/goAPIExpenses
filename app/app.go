package main

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/codehell/goAPIExpenses/dbo"
	"log"
	"github.com/codehell/goAPIExpenses/models"
	"database/sql"
	"time"
	"encoding/json"
)

var db *sql.DB

type ServeMux struct {
}

func (ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	path := url.Path

	parts := strings.Split(path, "/")
	fmt.Println(parts[0:], url.Path)
	w.Header().Set("Content-Type", "application/json")
	switch os := path;
		os {
	case "/expenses":
		expense := models.NewExpense(db)
		if err := expense.Get(34); err != nil {
			log.Fatal(err)
		}
		jsonExpense, err := json.Marshal(expense)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, string(jsonExpense))
	}

}

func main() {

	db = dbo.GetConnection()
	defer db.Close()

	/*expense := models.NewExpense(db)
	expense.Amount = 10.20
	expense.Description = "Cuncher"

	if err := expense.Create(); err != nil {
		log.Fatal(err)
	}*/

	s := &http.Server{
		Addr:           ":8080",
		Handler:        ServeMux{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
