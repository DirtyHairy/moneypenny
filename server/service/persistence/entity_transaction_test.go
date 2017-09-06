package persistence

import (
	"testing"
	"time"

	"github.com/dirtyhairy/moneypenny/server/model"
	"github.com/stretchr/testify/assert"
)

func assertTransactionCount(t *testing.T, p Provider, expected uint64) {
	transactionCount, err := p.CountTransactions()

	assert.NoError(t, err, "failed to query transaction count")
	assert.Equal(t, transactionCount, expected, "transaction count mismatch")
}

func TestInsertAndRetrieveTransaction(t *testing.T) {
	var err error

	p := createProvider(t)

	assertTransactionCount(t, p, 0)

	t0 := &model.Transaction{
		TransactionDate: time.Unix(time.Now().Unix(), 0),
		Amount:          666.,
		Title:           "foo",
		Description:     "bar",
	}

	assert.NoError(t, p.AddTransaction(t0), "insert failed")
	assertTransactionCount(t, p, 1)

	var t1 *model.Transaction
	t1, err = p.GetTransactionById(t0.Id)

	assert.NoError(t, err, "failed to retrieve transaction")

	assert.Equal(t, t0, t1, "transactions differ")
}

func TestGetAllTransactions(t *testing.T) {
	var err error

	p := createProvider(t)

	timestamp := time.Unix(time.Now().Unix(), 0)

	t0 := &model.Transaction{
		TransactionDate: timestamp.Add(time.Hour),
		Amount:          666,
		Title:           "hanni",
		Description:     "nanni",
	}

	t1 := &model.Transaction{
		TransactionDate: timestamp,
		Amount:          222,
		Title:           "foo",
		Description:     "bar",
	}

	assert.NoError(t, p.AddTransaction(t0), "insert failed")
	assert.NoError(t, p.AddTransaction(t1), "insert failed")

	assertTransactionCount(t, p, 2)

	var transactions []model.Transaction
	transactions, err = p.GetAllTransactions()

	assert.NoError(t, err, "failed to query all transactions")
	assert.Len(t, transactions, 2)
	assert.Equal(t, transactions[1], *t0)
	assert.Equal(t, transactions[0], *t1)
}

func TestDeleteTransaction(t *testing.T) {
	p := createProvider(t)

	assertTransactionCount(t, p, 0)

	t0 := &model.Transaction{
		TransactionDate: time.Unix(time.Now().Unix(), 0),
		Amount:          666.,
		Title:           "foo",
		Description:     "bar",
	}

	assert.NoError(t, p.AddTransaction(t0), "insert failed")
	assert.NoError(t, p.DeleteTransaction(t0), "delete failed")
	assert.Error(t, p.DeleteTransaction(t0), "deleting a non-existing transaction should fail")
}

func TestUpdateTransaction(t *testing.T) {
	var err error

	p := createProvider(t)

	assertTransactionCount(t, p, 0)

	t0 := &model.Transaction{
		TransactionDate: time.Unix(time.Now().Unix(), 0),
		Amount:          666.,
		Title:           "foo",
		Description:     "bar",
	}

	assert.NoError(t, p.AddTransaction(t0), "insert failed")

	t0.Amount = 333
	assert.NoError(t, p.UpdateTransaction(t0), "update failed")

	var t1 *model.Transaction
	t1, err = p.GetTransactionById(t0.Id)
	assert.NoError(t, err, "failed to retrieve transaction")

	assert.Equal(t, *t1, *t0, "transactions differ")
}