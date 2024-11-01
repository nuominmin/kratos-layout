package biz

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewService)

// Service .
type Service struct {
	c    *conf.Data
	repo Repo
	log  *log.Helper
}

// NewService .
func NewService(c *conf.Data, repo Repo, logger log.Logger) *Service {
	return &Service{c: c, repo: repo, log: log.NewHelper(logger)}
}

func (s *Service) Start(ctx context.Context) error {
	log.Info("starting service")
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	log.Info("stopping service")
	return nil
}
