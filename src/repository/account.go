package repository

import (
	"context"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type AccountRepo interface {
	FindAll(ctx context.Context) ([]entity.AccountEntity, error)
	FindByEmail(ctx context.Context, email string) (entity.AccountEntity, error)
	Insert(ctx context.Context, data entity.AccountEntity) (bool, error)
	Update(ctx context.Context, data entity.AccountEntity) (bool, error)
	DeleteById(ctx context.Context, id uint64) (bool, error)
}
