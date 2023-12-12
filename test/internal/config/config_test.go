package config

import (
	"fmt"
	"testing"

	"github.com/clibing/go-common/pkg/config"
	"github.com/clibing/go-common/pkg/model"
)

func TestShowConfig(t *testing.T) {

	type Config struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	}

	cfg := &Config{
		Host: "test.host.localhost",
		Port: 8080,
	}

	result, err := config.ShowConfig(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestLoadConfig(t *testing.T) {
	cfg := &model.Config{}
	err := config.LoadConfig[model.Config](cfg, "../../../configs/config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.Log.Level)

}
