package main

import (
	"database/sql"
	"github.com/codehell/goAPIExpenses/dbo"
	"github.com/codehell/goAPIExpenses/handler"
	"log"
	"net/http"
	"strings"
	"time"
	"io/ioutil"
	"fmt"
)

var db *sql.DB

type ServeMux struct {
}

func (ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	path := url.Path

	parts := strings.Split(path, "/")
	mainPath := parts[1]

	w.Header().Set("Content-Type", "application/json")

	switch os := mainPath; os {
	case "":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		index, err := ioutil.ReadFile("./vue/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprint(w, string(index))
	case "script":
		w.Header().Set("Content-Type", "application/javascript")
		index, err := ioutil.ReadFile("./vue/script.js")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprint(w, string(index))
	case "expenses":
		handler.ExpenseHandler(db, w, r, parts)
	}

}

func main() {

	db = dbo.GetConnection()
	defer db.Close()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        ServeMux{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
