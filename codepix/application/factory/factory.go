package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/luizcalazans16/imersao-full-cycle/codepix-go/application/usecase"
	"github.com/luizcalazans16/imersao-full-cycle/codepix-go/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixKeyRepository:      pixRepository,
	}

	return transactionUseCase
}
