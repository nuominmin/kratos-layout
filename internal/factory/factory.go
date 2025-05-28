package factory

import (
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/google/wire"
	"github.com/nuominmin/biz/captcha"
	"github.com/nuominmin/biz/krs/authorization/jwt"
	"github.com/nuominmin/biz/krs/authorization/token"
	"github.com/nuominmin/cache"
	workerreloader "github.com/nuominmin/worker-reloader"
)

var ProviderSet = wire.NewSet(
	NewCache,
	NewNewWorkerPool,
	NewToken,
	NewCaptcha,
	NewJwt,
)

func NewCache() cache.Cache {
	return cache.NewCache()
}

func NewNewWorkerPool() (workerreloader.WorkerPoolManager, func()) {
	wpm := workerreloader.NewWorkerPool()
	return wpm, wpm.StopAll
}

func NewToken() token.Service {
	return token.NewService()
}

func NewCaptcha() captcha.Service {
	return captcha.NewCaptcha()
}

func NewJwt(data *conf.Data) jwt.Service {
	return jwt.NewService([]byte(data.Env))
}
