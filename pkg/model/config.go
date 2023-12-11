package model

type Config struct {
	Log *Log `json:"log" yaml:"log" comment:"日志配置项"`
}

type Log struct {
	Level string `json:"level" yaml:"level" comment:"日志级别"` //日志级别
}
