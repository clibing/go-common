package config

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig[T any](value *T, path string) (err error) {
	if len(path) == 0 {
		err = fmt.Errorf("配置文件为空")
		return
	}

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
	_ = yaml.Unmarshal(data, value)
	return
}

/**
 * show config yml
 */
func ShowConfig[T any](input *T) (output string, err error) {
	var value bytes.Buffer
	encode := yaml.NewEncoder(&value)
	encode.SetIndent(2)
	err = encode.Encode(input)
	if err != nil {
		return
	}
	output = value.String()
	return
}
