package server

import (
	"io"

	"github.com/dirtyhairy/moneypenny/service/persistence"
)

type Config struct {
	Listen      string
	LogWriter   io.Writer
	Persistence persistence.Provider
	Debug       bool
}
