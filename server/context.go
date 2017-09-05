package server

import (
	"github.com/dirtyhairy/moneypenny/service/persistence"
	"github.com/labstack/echo"
)

type context struct {
	echo.Context
	persistence persistence.Provider
}

func (c *context) Persistence() persistence.Provider {
	return c.persistence
}
