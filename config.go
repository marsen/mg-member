package main

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type IConfig interface {
	Of() RootConfig
}

type RootConfig struct {
	DBUri string `env:"DBUri"`
}

type Config struct {
	vp       *viper.Viper
	instance *RootConfig
	once     sync.Once
}

func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("development")
	return &Config{vp: v}
}

func (cfg *Config) Of() *RootConfig {
	cfg.once.Do(func() {
		cfg.load()
	})
	return cfg.instance
}

func (cfg *Config) load() {
	// follow 12 factors principles https://www.12factor.net/config
	cfg.loadEnvironment()
	// just for development environment
	cfg.loadConfigFile()
	cfg.vp.Unmarshal(&cfg.instance)
	fmt.Printf("config instance :%v\n\n", cfg.instance)
}

func (cfg *Config) loadEnvironment() {
	cfg.vp.MustBindEnv("PORT")
	prefix := "Member"
	cfg.vp.BindEnv(prefix+".DB", "DB")
	cfg.vp.AutomaticEnv()
	cfg.vp.Unmarshal(&cfg.instance)
}

func (cfg *Config) loadConfigFile() {
	cfg.vp.SetConfigType("env")
	cfg.vp.AddConfigPath(".")
	cfg.vp.ReadInConfig()
}

// Mock this method is used to mock the config
// only for testing purpose
func (cfg *Config) Mock(v *viper.Viper) {
	cfg.vp = v
	cfg.vp.Unmarshal(&cfg.instance)
}

func (cfg *Config) SetConfigName(name string) {
	cfg.vp.SetConfigName(name)
}
