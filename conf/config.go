package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	orm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
	"github.com/infraboard/mcube/logger/zap"
)

const (
	CIPHER_TEXT_PREFIX = "@ciphered@"
)

// Config 应用配置
func newConfig() *Config {
	return &Config{
		App:     newDefaultAPP(),
		Log:     newDefaultLog(),
		MySQL:   newDefaultMySQL(),
		Cache:   newDefaultCache(),
		Mcenter: rpc.NewDefaultConfig(),
	}
}

type Config struct {
	App     *app        `toml:"app"`
	Log     *log        `toml:"log"`
	MySQL   *mysql      `toml:"mysql"`
	Mcenter *rpc.Config `toml:"mcenter"`
	Cache   *_cache     `toml:"cache"`
}

// InitGloabl 注入全局变量
func (c *Config) InitGloabl() error {
	// 加载全局配置单例
	global = c

	// 提前加载好 mcenter客户端
	err := rpc.LoadClientFromConfig(c.Mcenter)
	if err != nil {
		panic("load mcenter client from config error: " + err.Error())
	}
	return nil
}

// app 配置项
func newDefaultAPP() *app {
	return &app{
		Name:       "cmdb",
		HttpHost:   "127.0.0.1",
		HttpPort:   "8050",
		GRPCHost:   "127.0.0.1",
		GRPCPort:   "18050",
		EncryptKey: "defualt app encrypt key",
	}
}

type app struct {
	Name       string `toml:"name" env:"APP_NAME"`
	HttpHost   string `toml:"http_host" env:"APP_HTTP_HOST"`
	HttpPort   string `toml:"http_port" env:"APP_HTTP_PORT"`
	GRPCHost   string `toml:"grpc_host" env:"APP_GRPC_HOST"`
	GRPCPort   string `toml:"grpc_port" env:"APP_GRPC_PORT"`
	EncryptKey string `toml:"encrypt_key" env:"APP_ENCRYPT_KEY"`
}

func (a *app) HTTPAddr() string {
	return fmt.Sprintf("%s:%s", a.HttpHost, a.HttpPort)
}

func (a *app) GRPCAddr() string {
	return fmt.Sprintf("%s:%s", a.GRPCHost, a.GRPCPort)
}

// log 配置项
// newDefaultLog todo
func newDefaultLog() *log {
	return &log{
		Dir:    "logs",
		Level:  "debug",
		Format: "text",
		To:     "stdout",
	}
}

type log struct {
	Level  string    `toml:"level" env:"LOG_LEVEL"`
	Dir    string    `toml:"dir" env:"LOG_PATH_DIR"`
	Format LogFormat `toml:"format" env:"LOG_FORMAT"`
	To     LogTo     `toml:"to" env:"LOG_TO"`
}

// mysql 配置项
// newDefaultMySQL todo
func newDefaultMySQL() *mysql {
	return &mysql{
		Database:    "yCmdb",
		Host:        "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 200,
		MaxIdleConn: 50,
		MaxLifeTime: 1800,
		MaxIdleTime: 600,
	}
}

// MySQL todo
type mysql struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}

var (
	db *sql.DB
)

func (m *mysql) ORM() (*gorm.DB, error) {
	conn, err := m.GetDB()
	if err != nil {
		return nil, err
	}

	return gorm.Open(orm_mysql.New(orm_mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		// 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		PrepareStmt: true,
		// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。
		// 但这会降低性能，如果没有这方面的要求，可以在初始化时禁用它，这将获得大约 30%+ 性能提升
		SkipDefaultTransaction: true,
		// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法
		CreateBatchSize: 200,
	})
}

func (m *mysql) GetDB() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 加载全局数据库单例
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}

	return db, nil
}

// getDBConn use to get db connection pool
func (m *mysql) getDBConn() (*sql.DB, error) {
	zap.L().Named("config").Infof("connect to mysql: %s:%s", m.Host, m.Port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}

	return db, nil
}

func newDefaultCache() *_cache {
	return &_cache{
		Type:   "memory",
		Memory: memory.NewDefaultConfig(),
		Redis:  redis.NewDefaultConfig(),
	}
}

type _cache struct {
	Type   string         `toml:"type" json:"type" yaml:"type" env:"MCENTER_CACHE_TYPE"`
	Memory *memory.Config `toml:"memory" json:"memory" yaml:"memory"`
	Redis  *redis.Config  `toml:"redis" json:"redis" yaml:"redis"`
}

// func newDefaultKeyauth() *keyauth {
// 	return &keyauth{}
// }

// type keyauth struct {
// 	Host         string `toml:"host" env:"KEYAUTH_HOST"`
// 	Port         string `toml:"port" env:"KEYAUTH_PORT"`
// 	ClientID     string `toml:"client_id" env:"KEYAUTH_CLIENT_ID"`
// 	ClientSecret string `toml:"client_secret" env:"KEYAUTH_CLIENT_SECRET"`
// }

// func (a *keyauth) Addr() string {
// 	return a.Host + ":" + a.Port
// }

// func (a *keyauth) Client() (*kc.Client, error) {
// 	if kc.C() == nil {
// 		conf := kc.NewDefaultConfig()
// 		conf.SetAddress(a.Addr())
// 		zap.L().Infof("connect to keyauth: %s", a.Addr())
// 		conf.SetClientCredentials(a.ClientID, a.ClientSecret)
// 		client, err := kc.NewClient(conf)
// 		if err != nil {
// 			return nil, err
// 		}
// 		kc.SetGlobal(client)
// 	}

// 	return kc.C(), nil
// }
