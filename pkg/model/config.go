package model

type Config struct {
	Server *Server `json:"server" yaml:"server"`
	Log    *Log    `json:"log" yaml:"log" comment:"日志配置项"`
}

type Log struct {
	Level string `json:"level" yaml:"level" comment:"日志级别"` //日志级别
}

type Server struct {
	Host string `json:"host" yaml:"host" comment:""`
	Port int    `json:"port" yaml:"port" comment:""`
}
