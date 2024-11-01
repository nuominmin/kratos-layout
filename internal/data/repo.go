package data

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type repo struct {
	data *Data
	log  *log.Helper
}

// NewRepo .
func NewRepo(data *Data, logger log.Logger) biz.Repo {
	return &repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *repo) Save(ctx context.Context, g *biz.Entity) (*biz.Entity, error) {
	return g, nil
}
