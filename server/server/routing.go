package server

import (
	"net/http"

	"github.com/dirtyhairy/moneypenny/server/server/controller"
	"github.com/labstack/echo"
)

func (s *server) setupRouting() {
	s.echo.GET("/transaction", controller.GetAllTransactions)

	s.echo.Any("*", func(c echo.Context) error {
		c.NoContent(http.StatusNotFound)
		return nil
	})
}