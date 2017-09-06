package persistence

import (
	"errors"

	"github.com/dirtyhairy/moneypenny/server/model"
)

func (p *provider) GetTransactionById(id int) (user *model.Transaction, err error) {
	var result interface{}
	result, err = p.dbMap.Get(model.Transaction{}, id)

	if result != nil {
		user = result.(*model.Transaction)
	}

	return
}

func (p *provider) AddTransaction(transaction *model.Transaction) error {
	return p.dbMap.Insert(transaction)
}

func (p *provider) DeleteTransaction(transaction *model.Transaction) (err error) {
	var n int64
	n, err = p.dbMap.Delete(transaction)

	if err == nil && n == 0 {
		err = errors.New("entity not found")
	}

	return
}

func (p *provider) UpdateTransaction(transaction *model.Transaction) (err error) {
	var n int64
	n, err = p.dbMap.Update(transaction)

	if err == nil && n == 0 {
		err = errors.New("entity not found")
	}

	return
}

func (p *provider) CountTransactions() (count uint64, err error) {
	var result int64
	result, err = p.dbMap.SelectInt("SELECT COUNT(*) FROM `transaction`")

	count = uint64(result)

	return
}

func (p *provider) GetAllTransactions() (transactions []model.Transaction, err error) {
	_, err = p.dbMap.Select(&transactions, "SELECT * FROM `transaction` ORDER BY transaction_date  ASC")

	return
}