package config

import (
	"fmt"
)

type ServerListen struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}

func (s *ServerListen) ListenString() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *ServerListen) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type ServerConfig struct {
	HTTP ServerListen `mapstructure:"http"`
	GRPC ServerListen `mapstructure:"grpc"`
}
