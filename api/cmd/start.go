package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"github.com/ahwhy/yCmdb/api/protocol"
	"github.com/ahwhy/yCmdb/api/conf"
	"github.com/ahwhy/yCmdb/api/pkg"
	host "github.com/ahwhy/yCmdb/api/pkg/host/impl"
	searcher "github.com/ahwhy/yCmdb/api/pkg/resource/impl"
	syncer "github.com/ahwhy/yCmdb/api/pkg/syncer/impl"
)

type service struct {
	http *protocol.HTTPService
	log  logger.Logger
}

func newService(cnf *conf.Config) (*service, error) {
	http := protocol.NewHTTPService()
	svc := &service{
		http: http,
		log:  zap.L().Named("CLI"),
	}

	return svc, nil
}

func (s *service) start() error {
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
		// 初始化host.Service
		if err := host.Service.Config(); err != nil {
			return err
		}
		pkg.Host = host.Service

		// 初始化resource.Service
		if err := searcher.Service.Config(); err != nil {
			return err
		}
		pkg.Resource = searcher.Service

		//  初始化Syncer.Service
		if err := syncer.Service.Config(); err != nil {
			return err
		}
		pkg.Syncer = syncer.Service

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

func init() {
	RootCmd.AddCommand(serviceCmd)
}
