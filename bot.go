package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
	"time"
)

type Bot struct {
	tgBot *tgbotapi.BotAPI

	id2FlomoAPI *sync.Map
}

func NewBot(token string) (*Bot, error) {
	b := &Bot{}
	b.id2FlomoAPI = &sync.Map{}
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
	if update.FromChat().Type != "private" {
		// 不是私聊的消息不处理
		return
	}

	if update.Message.IsCommand() {
		b.handleCommand(update)
		return
	}

	userID := update.Message.From.ID
	chatID := update.Message.Chat.ID
	msg := update.Message.Text
	note := &FlomoNote{Content: msg}

	token, ok := b.getFlomoToken(userID)
	if !ok {
		b.auth(chatID)
		return
	}
	err := PostFlomoNote(token, note)
	if err != nil {
		b.fail(chatID, err)
		return
	}
	b.send(chatID, SuccessTip)
}

func (b *Bot) getFlomoToken(id int64) (string, bool) {
	val, ok := b.id2FlomoAPI.Load(id)
	if !ok {
		return "", false
	}

	return val.(string), true
}

func (b *Bot) send(chatID int64, content string) {
	log.Println("new note send send")
	msg := tgbotapi.NewMessage(chatID, content)
	b.tgBot.Send(msg)
}

func (b *Bot) fail(chatID int64, e error) {
	log.Println("new note send fail")
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s: %s", FailTip, e.Error()))
	b.tgBot.Send(msg)
}

func (b *Bot) auth(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, AuthTip)
	b.tgBot.Send(msg)
}

func (b *Bot) handleCommand(update tgbotapi.Update) {
	cmd := update.Message.Command()

	switch cmd {
	case "ping":
		b.send(update.Message.Chat.ID, PingTip)
	case "start":
		b.send(update.Message.Chat.ID, StartTip)
	case "bind":
		b.bind(update)
	case "unbind":
		b.unBind(update)
	default:
		return
	}
}

func (b *Bot) bind(update tgbotapi.Update) {
	token := update.Message.CommandArguments()
	if !validFlomoAPI(token) {
		b.send(update.Message.Chat.ID, InvalidFlomoAPITip)
		return
	}
	userID := update.Message.From.ID
	b.id2FlomoAPI.Store(userID, token)
	b.send(update.Message.Chat.ID, BindTip)
}

func (b *Bot) unBind(update tgbotapi.Update) {
	userID := update.Message.From.ID
	b.id2FlomoAPI.Delete(userID)
	b.send(update.Message.Chat.ID, UnBindTip)
}
