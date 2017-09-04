package persistence

import (
	"github.com/dirtyhairy/moneypenny/model"
)

type Provider interface {
	GetMeta() (*model.Meta, error)
	Close() error
}
