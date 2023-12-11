package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"github.com/clibing/go-common/internal/pkg/model"
)

func LoadConfig(path string) (config *model.Config, err error) {
	if len(path) > 0 {
		var data []byte
		data, err = os.ReadFile(path)
		if err != nil {
			err = fmt.Errorf("加载配置异常: %s", err.Error())
			return
		}
		if len(data) == 0 {
			err = fmt.Errorf("配置文件为空")
			return
		}
		config = &model.Config{}
		_ = yaml.Unmarshal(data, config)
		return
	}
	err = fmt.Errorf("配置文件为空")
	return
}

/**
 * show config yml
 */
func ShowConfig(config *model.Config) error {
	cfg, err := yaml.Marshal(config)
	if err != nil {
		logrus.Info("show config yaml")
		return err
	}
	// logrus.Infof("%s", cfg)
	fmt.Printf("%s\n", cfg)
	return nil
}
