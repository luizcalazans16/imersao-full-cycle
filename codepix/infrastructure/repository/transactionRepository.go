package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/luizcalazans16/imersao-full-cycle/codepix/domain/model"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (repo *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := repo.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := repo.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	repo.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &transaction, nil
}
