package server

import (
	pb "github.com/go-kratos/kratos-layout/api"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	nethttp "net/http"
	"time"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, svc *service.Service) *http.Server {
	var opts = []http.ServerOption{
		http.Timeout(time.Second * 30),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "User-Agent", "Content-Length", "Access-Control-Allow-Credentials"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}), handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowCredentials(),
		)),
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(nethttp.StatusOK)
	})
	pb.RegisterV1HTTPServer(srv, svc)
	return srv
}
