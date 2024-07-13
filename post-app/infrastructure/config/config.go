package config

import (
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/logger"

	"github.com/caarlos0/env/v11"
)

type AppEnvironment string

const (
	Local       AppEnvironment = "local"
	Development AppEnvironment = "development"
	Staging     AppEnvironment = "staging"
	Production  AppEnvironment = "production"
)

type Config struct {
	App AppConfig `envPrefix:"APP_"`
	DB  DBConfig  `envPrefix:"DB_"`
}

type AppConfig struct {
	AppHost        string         `env:"HOST" envDefault:"0.0.0.0"`
	AppPort        uint           `env:"PORT" envDefault:"8000"`
	AppEnvironment AppEnvironment `env:"ENVIRONMENT" envDefault:"local"`
}

type DBConfig struct {
	Host     string `env:"HOST,required"`
	Port     uint   `env:"PORT,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Name     string `env:"NAME,required"`
	SSLMode  bool   `env:"SSL_MODE" envDefault:"false"`
	// DB接続開始時のタイムアウト時間(秒)
	Timeout uint `env:"TIMEOUT" envDefault:"10"`

	// 最大コネクション数
	MaxConnections uint `env:"MAX_OPEN_CONNECTIONS" envDefault:"10"`
	// 最小コネクション数
	MinConnections uint `env:"MIN_OPEN_CONNECTIONS" envDefault:"5"`
	// コネクションを利用可能な最長時間(秒)
	ConnectionMaxLifetime uint `env:"CONNECTION_MAX_LIFETIME" envDefault:"300"`
	// コネクションがアイドル状態で保持される時間(秒)
	ConnectionMaxIdletime uint `env:"CONNECTION_MAX_IDLE_TIME" envDefault:"300"`
}

func NewAppConfig() (*AppConfig, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		logger.Error("アプリケーション設定の読み込みに失敗しました", "err", err)
		return nil, err
	}
	return &cfg.App, nil
}

func NewDBConfig() (*DBConfig, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		logger.Error("DB設定の読み込みに失敗しました", "err", err)
		return nil, err
	}
	return &cfg.DB, nil
}
