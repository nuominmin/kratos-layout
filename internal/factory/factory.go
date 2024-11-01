package factory

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/google/wire"
	"github.com/nuominmin/cache"
	"github.com/nuominmin/notifier"
	"github.com/nuominmin/notifier/lark"
	workerreloader "github.com/nuominmin/worker-reloader"
)

var ProviderSet = wire.NewSet(
	NewAlert,
	NewCache,
	NewNewWorkerPool,
	NewEthClient,
)

func NewAlert(data *conf.Data) (notifier.Notifier, error) {
	alert := lark.NewNotifier(data.GetAlertTokens()...)
	if env := data.GetEnv(); env != "" {
		alert.SetIdentity(env)
	}
	return alert, nil
}

func NewCache() cache.Cache {
	return cache.NewCache()
}

func NewNewWorkerPool() (workerreloader.WorkerPoolManager, func()) {
	wpm := workerreloader.NewWorkerPool()
	return wpm, wpm.StopAll
}

func NewEthClient(data *conf.Data) (*ethclient.Client, error) {
	rpcClient, err := rpc.Dial(data.EthereumEndpoint)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(rpcClient), nil
}
