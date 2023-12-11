package model

type Config struct {
	Server *Server `json:"server" yaml:"server" comment:"服务端配置项"`
	Log    *Log    `json:"log" yaml:"log" comment:"日志配置项"`
}
type Server struct {
	Host string `json:"host" yaml:"host" comment:"监听地址"` // gin启动监听的端口
	Port int    `json:"port" yaml:"port" comment:"监听端口"` // gin启动监听的端口
}

type Log struct {
	Level string `json:"level" yaml:"level" comment:"日志级别"` //日志级别
}
