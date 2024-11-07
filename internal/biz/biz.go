package biz

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nuominmin/notifier"
	workerreloader "github.com/nuominmin/worker-reloader"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewService)

// Service .
type Service struct {
	alert      notifier.Notifier
	workerPool workerreloader.WorkerPoolManager
	c          *conf.Data
	repo       Repo
	log        *log.Helper
}

// NewService .
func NewService(c *conf.Data, alert notifier.Notifier, workerPool workerreloader.WorkerPoolManager, repo Repo, logger log.Logger) *Service {
	return &Service{
		alert:      alert,
		workerPool: workerPool,
		c:          c,
		repo:       repo,
		log:        log.NewHelper(logger),
	}
}

func (s *Service) Start(ctx context.Context) error {
	log.Info("starting service")
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	log.Info("stopping service")
	return nil
}
