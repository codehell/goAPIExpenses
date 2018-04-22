package models

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

type expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (e *expense) Create() error {

	q := `INSERT INTO expenses (amount, description)
			VALUES ($1, $2) 
			RETURNING id, amount, description, created_at, updated_at`
	row, err := db.Query(q, e.Amount, e.Description)
	defer row.Close()
	if err != nil {
		return err
	}

	row.Next()
	if err = row.Scan(&e.ID, &e.Amount, &e.Description, &e.CreatedAt, &e.UpdatedAt); err != nil {
		return err
	}
	if err = row.Err(); err != nil {
		return err
	}

	fmt.Println(e)

	return nil
}

func (e *expense) Get(id int64) error {
	q := `SELECT * FROM expenses WHERE id = $1`
	rows, err := db.Query(q, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&e.ID, &e.Amount, &e.Description, &e.CreatedAt, &e.UpdatedAt)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (e expense) All() ([]expense, error) {

	var expenses []expense

	q := `SELECT id, amount, description, created_at, updated_at FROM expenses`

	rows, err := db.Query(q)
	defer rows.Close()
	if err != nil {
		return expenses, err
	}
	for rows.Next() {
		rows.Scan(&e.ID, &e.Amount, &e.Description, &e.CreatedAt, &e.UpdatedAt)
		expenses = append(expenses, e)
	}
	if err = rows.Err(); err != nil {
		return expenses, err
	}
	return expenses, nil
}

func prepare(q string) (*sql.Stmt, error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func NewExpense(database *sql.DB) *expense {
	db = database
	return new(expense)
}
