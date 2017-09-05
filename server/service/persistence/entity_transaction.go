package persistence

import "github.com/dirtyhairy/moneypenny/server/model"

func (p *provider) GetTransactionById(id uint64) (user *model.Transaction, err error) {
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
	_, err = p.dbMap.Delete(transaction)

	return
}

func (p *provider) CountTransactions() (count uint64, err error) {
	var result int64
	result, err = p.dbMap.SelectInt("SELECT COUNT(*) FROM `transaction`")

	count = uint64(result)

	return
}
