package main

import (
	"errors"
	"log"
	"os"
)

var (
	SuccessTip = "🟢 发送成功"

	FailTip = "🔴 发送失败"

	StartTip = "🤖 欢迎使用Flomo机器人"

	PingTip = "🏓 Pong!"
)

var (
	FlomoToken string
	TgBotToken string
)

var (
	FlomoEnvKey = "FLOMO_API_URL"

	TgBotEnvKey = "TG_BOT_TOKEN"
)

func retrieveToken() error {
	// 从环境变量中检查
	apiUrl := os.Getenv(FlomoEnvKey)
	botToken := os.Getenv(TgBotEnvKey)

	if apiUrl != "" && botToken != "" {
		FlomoToken = apiUrl
		TgBotToken = botToken
		return nil
	}

	// 从配置文件中检查
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	apiUrl = cfg.FlomoAPI
	botToken = cfg.TgBotToken
	if apiUrl != "" && botToken != "" {
		FlomoToken = apiUrl
		TgBotToken = botToken
		return nil
	}
	return errors.New("flomo API URL and tg bot token not correct")
}
