package factory

import "github.com/victordev13/fullcycle5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
