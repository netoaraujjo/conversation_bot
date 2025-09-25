package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	TEXT int64 = iota
	PHOTO
)

type MatcherFunc func(update tgbotapi.Update) bool
type HandlerFunc func(bot *tgbotapi.BotAPI, update tgbotapi.Update) int

// Define o formato das funções que tratarão os eventos: mensagem, callback query, comandos
// Retorna o próximo estado
type EventHandler struct {
	Match   MatcherFunc
	Handler HandlerFunc
}

func MessageHandler(msgType int64) MatcherFunc {
	return func(update tgbotapi.Update) bool {
		if update.Message != nil && !update.Message.IsCommand() {
			if (update.Message.Text != "" && msgType == TEXT) || (len(update.Message.Photo) > 0 && msgType == PHOTO) {
				return true
			}
		}
		return false
	}
}

func CommandHandler(command string) MatcherFunc {
	return func(update tgbotapi.Update) bool {
		return update.Message != nil && update.Message.IsCommand() && command == update.Message.Command()
	}
}

func CallbackQueryHandler(data string) MatcherFunc {
	return func(update tgbotapi.Update) bool {
		return update.CallbackQuery != nil && update.CallbackData() == data
	}
}

type ConversationHandler struct {
	EntryPoints []EventHandler
	States      map[int64][]EventHandler
	Fallbacks   []EventHandler
}

func NewConversationHandler() *ConversationHandler {
	return &ConversationHandler{
		States: make(map[int64][]EventHandler),
	}
}
