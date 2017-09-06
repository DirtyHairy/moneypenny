package server

import (
	"net/http"

	"github.com/dirtyhairy/moneypenny/server/server/controller"
	"github.com/labstack/echo"
)

func (s *server) setupRouting() {
	s.echo.GET("/transactions", controller.GetAllTransactions)
	s.echo.POST("/transactions", controller.AddTransaction)
	s.echo.GET("/transactions/:id", controller.GetTransactionById)
	s.echo.DELETE("/transactions/:id", controller.DeleteTransactionById)
	s.echo.PUT("/transactions/:id", controller.UpdateTransactionById)

	s.echo.Any("*", func(c echo.Context) error {
		c.NoContent(http.StatusNotFound)
		return nil
	})
}