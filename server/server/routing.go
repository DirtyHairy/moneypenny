package server

import (
	"net/http"

	"github.com/dirtyhairy/moneypenny/server/server/controller"
	"github.com/labstack/echo"
)

func (s *server) setupRouting() {
	notFound := func(c echo.Context) error {
		c.NoContent(http.StatusNotFound)
		return nil
	}

	s.echo.GET("/api/transactions", controller.GetAllTransactions)
	s.echo.POST("/api/transactions", controller.AddTransaction)
	s.echo.GET("/api/transactions/:id", controller.GetTransactionById)
	s.echo.DELETE("/api/transactions/:id", controller.DeleteTransactionById)
	s.echo.PUT("/api/transactions/:id", controller.UpdateTransactionById)

	if s.config.StaticFS != nil {
		s.echo.Any("/api/*", notFound)
		s.echo.GET("*", echo.WrapHandler(http.FileServer(s.config.StaticFS)))
	} else {
		s.echo.Any("*", notFound)
	}
}