package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

func retrieveToken() error {
	// 从环境变量中检查
	botToken := os.Getenv(TgBotEnvKey)

	if botToken != "" {
		TgBotToken = botToken
		return nil
	}

	// 从配置文件中检查
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	botToken = cfg.TgBotToken
	if botToken != "" {
		TgBotToken = botToken
		return nil
	}
	return errors.New("tg bot token not correct")
}

func validFlomoAPI(url string) bool {
	// 检查 URL 是否为空
	if url == "" {
		return false
	}

	// 检查 URL 是否以 https://flomoapp.com/iwh/ 开头
	if !strings.HasPrefix(url, "https://flomoapp.com/iwh/") {
		return false
	}

	// 考虑正则
	return true
}
