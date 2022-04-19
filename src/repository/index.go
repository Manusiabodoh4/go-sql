package repository

import (
	"context"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type Repository interface {
	Find(ctx context.Context, channel chan entity.TemplateChannelResponse)
	FindWithParam(ctx context.Context, channel chan entity.TemplateChannelResponse, param string, args ...interface{})
	InsertOne(ctx context.Context, channel chan entity.TemplateChannelResponse, args ...interface{})
	InsertMany(ctx context.Context, channel chan entity.TemplateChannelResponse, value []map[string]interface{})
	Update(ctx context.Context, channel chan entity.TemplateChannelResponse, set string, param string, args ...interface{})
	Delete(ctx context.Context, channel chan entity.TemplateChannelResponse, param string, args ...interface{})
}
