package model

import (
	"time"
)

type Transaction struct {
	Id              int       `db:"id,primarykey,autoincrement" json:"id"`
	TransactionDate time.Time `db:"transaction_date,notnull" json:"transactionDate"`
	Amount          uint64    `db:"amount,notnull" json:"amount"`
	Title           string    `db:"title,notnull" json:"title"`
	Description     string    `db:"description,notnull" json:"description"`
}