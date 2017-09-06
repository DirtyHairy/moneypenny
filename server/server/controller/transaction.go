package controller

import (
	"net/http"

	"github.com/dirtyhairy/moneypenny/server/model"
	"github.com/dirtyhairy/moneypenny/server/server/context"

	"github.com/labstack/echo"
)

func GetAllTransactions(c echo.Context) (err error) {
	persistence := c.(context.Context).Persistence()

	var transactions []model.Transaction

	if transactions, err = persistence.GetAllTransactions(); err == nil {
		c.JSON(http.StatusOK, transactions)
	}

	return
}