package usecase

import (
	"errors"
	"log"

	"github.com/luizcalazans16/imersao-full-cycle/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepository
	PixKeyRepository      model.PixKeyRepository
}

func (t *TransactionUseCase) Register(accountId string, amount float64, pixKeyto string, pixKeyKindTo string, description string, id string) (*model.Transaction, error) {
	account, err := t.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixKeyRepository.FindKeyByKind(pixKeyto, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description, id)
	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)
	if transaction.Base.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("Unable to process this transaction")

}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
