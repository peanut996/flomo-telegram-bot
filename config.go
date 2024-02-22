package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

// LoadConfig 函数用于加载配置文件并解析其中的配置项。
func LoadConfig(filename string) (*Config, error) {
	// 读取配置文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解析配置文件内容
	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Config 结构体用于表示配置文件中的配置项。
type Config struct {
	FlomoAPI   string `yaml:"flomo_api"`
	TgBotToken string `yaml:"tgbot_token"`
}
