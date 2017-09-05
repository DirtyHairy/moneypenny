package model

import (
	"time"
)

type Transaction struct {
	Id              uint64    `db:"id,primarykey,autoincrement"`
	TransactionDate time.Time `db:"transaction_date,notnull"`
	Amount          uint64    `db:"amount,notnull"`
	Title           string    `db:"title,notnull"`
	Description     string    `db:"description,notnull"`
}
