package repository

import (
	"context"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type AccountRepo interface {
	FindAll(ctx context.Context) ([]entity.AccountEntity, error)
	FindWithParam(ctx context.Context, data entity.AccountEntity) (entity.AccountEntity, error)
	Insert(ctx context.Context, data entity.AccountEntity) (bool, error)
	Update(ctx context.Context, data entity.AccountEntity) (bool, error)
	Delete(ctx context.Context, data entity.AccountEntity) (bool, error)
}
