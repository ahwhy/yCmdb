package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ahwhy/yCmdb/app"
	"github.com/ahwhy/yCmdb/conf"
	"github.com/ahwhy/yCmdb/protocol"

	// 注册所有服务
	_ "github.com/ahwhy/yCmdb/app/all"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "yCmdb后端API服务",
	Long:  `yCmdb后端API服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局日志配置
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// Ioc初始化
		// 初始化全局app
		if err := app.InitAllApp(); err != nil {
			return err
		}

		// 启动服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 初始化服务
		svc, err := newService(conf.C())
		if err != nil {
			return err
		}

		// 等待信号处理
		go svc.waitSign(ch)
		// 启动服务
		if err := svc.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}

		return nil
	},
}

// config 为全局变量, 只需要load即可全局使用
func loadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	case "etcd":
		return errors.New("not implemented")
	default:
		return errors.New("unknown config type")
	}

	return nil
}

func newService(cnf *conf.Config) (*service, error) {
	http := protocol.NewHTTPService()
	grpc := protocol.NewGRPCService()
	svc := &service{
		http: http,
		grpc: grpc,
		log:  zap.L().Named("CLI"),
	}

	return svc, nil
}

type service struct {
	http *protocol.HTTPService
	grpc *protocol.GRPCService
	log  logger.Logger
}

func (s *service) start() error {
	s.log.Infof("loaded grpc app: %s", app.LoadedGrpcApp())
	s.log.Infof("loaded http app: %s", app.LoadedHttpApp())
	s.log.Infof("loaded internal app: %s", app.LoadedInternalApp())

	go s.grpc.Start()
	return s.http.Start()
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		default:
			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
			if err := s.http.Stop(); err != nil {
				s.log.Errorf("graceful shutdown err: %s, force exit", err)
			}
			s.log.Infof("service stop complete")
			return
		}
	}
}

// log 为全局变量, 只需要load 即可全局使用, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	lc := conf.C().Log
	lv, err := zap.NewLevel(lc.Level)

	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	zapConfig.Files.RotateOnStartup = false

	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)

	return nil
}

func init() {
	RootCmd.AddCommand(serviceCmd)
}
