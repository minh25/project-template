package config

import "strconv"

type Listen struct {
	Port uint16 `mapstructure:"port"`
}

func (l *Listen) String() string {
	return strconv.FormatUint(uint64(l.Port), 10)
}

type ServerConfig struct {
	Http Listen `mapstructure:"http"`
	Grpc Listen `mapstructure:"grpc"`
}
