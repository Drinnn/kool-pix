package factory

import (
	"github.com/Drinnn/kool-pix/application/usecase"
	"github.com/Drinnn/kool-pix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUsecaseFactory(database *gorm.DB) usecase.TransactionUsecase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUsecase := usecase.TransactionUsecase{
		TransactionRepository: &transactionRepository,
		PixRepository:         &pixRepository,
	}

	return transactionUsecase
}
