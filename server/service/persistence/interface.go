package persistence

import (
	"github.com/dirtyhairy/moneypenny/server/model"
)

type Provider interface {
	Close() error

	GetMeta() (*model.Meta, error)

	AddTransaction(*model.Transaction) error
	GetTransactionById(id uint64) (*model.Transaction, error)
	DeleteTransaction(*model.Transaction) error
	CountTransactions() (uint64, error)
}
