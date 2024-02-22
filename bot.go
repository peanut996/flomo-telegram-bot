package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type Bot struct {
	tgBot *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	b := &Bot{}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	b.tgBot = bot
	return b, nil
}
func (b *Bot) Run() {
	b.fetchUpdates()
}

func (b *Bot) fetchUpdates() {
	config := tgbotapi.NewUpdate(0)
	config.Timeout = 60
	config.AllowedUpdates = []string{"message", "edited_message", "channel_post", "edited_channel_post", "chat_member"}

	botChannel := b.tgBot.GetUpdatesChan(config)
	for {
		select {
		case update, ok := <-botChannel:
			if !ok {
				b.tgBot.StopReceivingUpdates()
				botChannel = b.tgBot.GetUpdatesChan(config)
				log.Println("[FetchUpdates] channel closed, fetch again")
				continue
			}
			go b.handleUpdate(update)
		case <-time.After(30 * time.Second):
		}
	}
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		b.handleCommand(update)
		return
	}
	chatID := update.Message.Chat.ID
	msg := update.Message.Text

	note := &FlomoNote{Content: msg}

	err := PostFlomoNote(FlomoToken, note)
	if err != nil {
		// error handle
		b.fail(chatID, err)
		return
	}
	// success
	b.success(chatID)
}

func (b *Bot) success(chatID int64) {
	log.Println("new note send success")
	msg := tgbotapi.NewMessage(chatID, SuccessTip)
	b.tgBot.Send(msg)
}

func (b *Bot) fail(chatID int64, e error) {
	log.Println("new note send fail")
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s: %s", FailTip, e.Error()))
	b.tgBot.Send(msg)
}

func (b *Bot) hello(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, StartTip)
	b.tgBot.Send(msg)
}

func (b *Bot) pong(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, PingTip)
	b.tgBot.Send(msg)
}

func (b *Bot) handleCommand(update tgbotapi.Update) {
	cmd := update.Message.Command()

	switch cmd {
	case "ping":
		b.pong(update.Message.Chat.ID)
	case "start":
		b.hello(update.Message.Chat.ID)
	default:
		return
	}
}
