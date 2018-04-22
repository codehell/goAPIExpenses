package models

import (
	"database/sql"
	"time"
	"fmt"
)

var db *sql.DB

type expense struct {
	ID          int
	Amount      float64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

func (expense) All() string {
	return "Ahi va todo"
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
