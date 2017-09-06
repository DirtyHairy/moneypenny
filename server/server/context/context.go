package context

import (
	"github.com/dirtyhairy/moneypenny/server/service/persistence"
	"github.com/labstack/echo"
)

type Context interface {
	echo.Context

	Persistence() persistence.Provider
}