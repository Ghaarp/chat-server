package config

import (
	"fmt"
	"os"
)

const (
	chatHostEnvName = "CHAT_HOST"
	chatPortEnvName = "CHAT_PORT"
)

type ChatConfig interface {
	Address() string
}

type chatConfig struct {
	host string
	port string
}

func (cc *chatConfig) Address() string {
	return fmt.Sprintf(":%s", cc.port)
}

func NewChatConfig() (ChatConfig, error) {
	host := os.Getenv(chatHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("No host name in .env")
	}

	port := os.Getenv(chatPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("No port in .env")
	}

	config := chatConfig{
		host: host,
		port: port,
	}

	return &config, nil
}
