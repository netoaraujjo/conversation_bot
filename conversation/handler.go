package conversation

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Define o formato das funções que tratarão os eventos: mensagem, callback query, comandos
// Retorna o próximo estado
type EventHandler func(bot *tgbotapi.BotAPI, update tgbotapi.Update) int

type ConversationHandler struct {
}
