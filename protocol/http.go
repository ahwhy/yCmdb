package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/conf"

	"github.com/infraboard/mcube/http/middleware/accesslog"
	"github.com/infraboard/mcube/http/middleware/cors"
	"github.com/infraboard/mcube/http/middleware/recovery"
	"github.com/infraboard/mcube/http/router"
	"github.com/infraboard/mcube/http/router/httprouter"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {

	r := httprouter.New()
	r.Use(recovery.NewWithLogger(zap.L().Named("Recovery")))
	r.Use(accesslog.NewWithLogger(zap.L().Named("AccessLog")))
	r.Use(cors.AllowAll())
	r.EnableAPIRoot()
	// r.SetAuther(auther)
	r.Auth(true)

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTPAddr(),
		Handler:           cors.AllowAll().Handler(r),
	}

	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

// HTTPService http服务
type HTTPService struct {
	r      router.Router
	l      logger.Logger
	c      *conf.Config
	server *http.Server
}

func (s *HTTPService) Addr() string {
	return fmt.Sprintf("%s/api/v1", s.c.App.Name)
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 配置子服务路由
	app.LoadHttpApp(s.Addr(), s.r)
	
	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}

	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}

	return nil
}
