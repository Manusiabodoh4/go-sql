package repository

import (
	"context"
	"sync"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type AccountRepo interface {
	FindAll(ctx context.Context, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse)
	FindByEmail(ctx context.Context, email string, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse)
	Insert(ctx context.Context, data entity.AccountEntity, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse)
	Update(ctx context.Context, data entity.AccountEntity, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse)
	DeleteById(ctx context.Context, email string, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse)
}
