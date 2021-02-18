package cfg

import "fmt"

type Config struct {
	Host string
	Port int
}

func (config Config) SocketAddress() string {
	return config.Host + ":" + fmt.Sprint(config.Port)
}
