package models

import "database/sql"

type model struct {
	DB *sql.DB
	Model interface{}
}

func NewModel(database *sql.DB, concreteModel interface{}) *model {
	model := model{
		DB: database,
		Model: concreteModel,
	}
	return &model
}
