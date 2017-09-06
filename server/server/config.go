package server

import (
	"io"
	"net/http"

	"github.com/dirtyhairy/moneypenny/server/service/persistence"
)

type Config struct {
	Listen      string
	LogWriter   io.Writer
	Persistence persistence.Provider
	Debug       bool
	StaticFS    http.FileSystem
}