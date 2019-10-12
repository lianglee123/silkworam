package config

import (
	"sync"
	"github.com/lianglee123/easyconfig"
	"log"
)

type DemoConfig struct {
	Debug bool

	DB struct {
		Driver string
		URL string
		LogSQL bool
		TableNamePrefix string
		MaxIdleConns int
		MaxOpenConns int
	}

	Redis struct {
		Host string
		Port string
		Pwd string `config:"default:5170"`
	}

	Logger struct {
		Level string
		Formatter string
	}

	Http struct {
		Host string
		Port string
	}

	GRPC struct {
		Host string
		Port string
	}
}


var (
	conf *DemoConfig
	once sync.Once
)


func Load() *DemoConfig {
	once.Do(func() {
		opt := &easyconfig.LoadOption{
			EnvPrefix: "DEMO",
			ConfigFilePath: "./config.yaml",
		}
		err := easyconfig.LoadConfig(conf, opt)
		if err != nil {
			log.Fatalf("load config fail: %s", err.Error())
		}
	})
	return conf
}
