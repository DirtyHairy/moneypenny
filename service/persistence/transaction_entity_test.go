package persistence

import (
	"testing"
	"time"

	"github.com/dirtyhairy/moneypenny/model"
)

func assertTransactionCount(t *testing.T, p Provider, expected uint64) {
	transactionCount, err := p.CountTransactions()

	failIfError(t, err, "failed to query transaction count: %s", err)
	if transactionCount != expected {
		t.Fatalf("transaction count wrong: got %d, expected %d", transactionCount, expected)
	}
}

func assertTransactionIdentitiy(t *testing.T, t0, t1 *model.Transaction) {
	if t0.TransactionDate.Unix() != t1.TransactionDate.Unix() {
		t.Fatalf(
			"transaction time mismatch; expected %s, got %s",
			t0.TransactionDate.String(),
			t1.TransactionDate.String(),
		)
	}

	if t0.Amount != t1.Amount {
		t.Fatalf(
			"transaction amount mismatch; expected %v, got %v",
			t0.Amount,
			t1.Amount,
		)
	}

	if t0.Title != t1.Title {
		t.Fatalf(
			"transaction title mismatch; expected %v, got %v",
			t0.Title,
			t1.Title,
		)
	}

	if t0.Description != t1.Description {
		t.Fatalf(
			"transaction description mismatch; expected %v, got %v",
			t0.Description,
			t1.Description,
		)
	}
}

func TestInsertAndRetrieveTransaction(t *testing.T) {
	var err error

	p := createProvider(t)

	assertTransactionCount(t, p, 0)

	t0 := &model.Transaction{
		TransactionDate: time.Now(),
		Amount:          666.,
		Title:           "foo",
		Description:     "bar",
	}

	failIfError(t, p.AddTransaction(t0), "insert failed: %s", err)

	assertTransactionCount(t, p, 1)

	var t1 *model.Transaction
	t1, err = p.GetTransactionById(t0.Id)
	failIfError(t, err, "failed to retrieve transaction: %s", err)

	assertTransactionIdentitiy(t, t0, t1)
}
