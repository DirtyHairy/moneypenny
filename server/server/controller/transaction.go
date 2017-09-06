package controller

import (
	"net/http"
	"strconv"
	"time"

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

func AddTransaction(c echo.Context) (err error) {
	persistence := c.(context.Context).Persistence()

	transaction := &model.Transaction{
		TransactionDate: time.Now(),
	}

	if err = c.Bind(transaction); err != nil {
		return
	}

	if err = persistence.AddTransaction(transaction); err != nil {
		return
	}

	c.JSON(http.StatusOK, transaction)

	return
}

func retrieveTransaction(c echo.Context) (transaction *model.Transaction, err error) {
	persistence := c.(context.Context).Persistence()

	var id int
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.NoContent(http.StatusBadRequest)
		return
	}

	if transaction, err = persistence.GetTransactionById(id); err != nil {
		return
	}

	if transaction == nil {
		c.NoContent(http.StatusNotFound)
		return
	}

	return
}

func GetTransactionById(c echo.Context) (err error) {
	var transaction *model.Transaction
	transaction, err = retrieveTransaction(c)

	if transaction != nil || err != nil {
		c.JSON(http.StatusOK, transaction)
	}

	return
}

func DeleteTransactionById(c echo.Context) (err error) {
	var transaction *model.Transaction
	transaction, err = retrieveTransaction(c)

	if transaction == nil || err != nil {
		return
	}

	persistence := c.(context.Context).Persistence()

	if err = persistence.DeleteTransaction(transaction); err != nil {
		return
	}

	c.NoContent(http.StatusNoContent)

	return
}

func UpdateTransactionById(c echo.Context) (err error) {
	incomingTransaction := &model.Transaction{
		TransactionDate: time.Now(),
	}

	if err = c.Bind(incomingTransaction); err != nil {
		return
	}

	var transaction *model.Transaction
	transaction, err = retrieveTransaction(c)

	if transaction == nil || err != nil {
		return
	}

	persistence := c.(context.Context).Persistence()

	incomingTransaction.Id = transaction.Id
	if err = persistence.UpdateTransaction(incomingTransaction); err != nil {
		return
	}

	c.JSON(http.StatusOK, incomingTransaction)

	return
}