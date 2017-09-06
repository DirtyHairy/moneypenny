package persistence

import (
	"github.com/dirtyhairy/moneypenny/server/model"
)

type Provider interface {
	Close() error

	GetMeta() (*model.Meta, error)

	GetAllTransactions() ([]model.Transaction, error)
	AddTransaction(*model.Transaction) error
	UpdateTransaction(*model.Transaction) error
	GetTransactionById(id int) (*model.Transaction, error)
	DeleteTransaction(*model.Transaction) error
	CountTransactions() (uint64, error)
}