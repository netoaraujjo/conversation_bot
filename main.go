package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Usuario struct {
	State      int
	ID         int
	Name       string
	Username   string
	Parameters map[string]string
}

func main() {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %s", err)
	}
	botToken := os.Getenv("BOT_TOKEN")
	// lupaAPIKey := os.Getenv("LUPA_API_KEY")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Erro ao inicializar o bot: %s", err)
	}
	// bot.Debug = true
	log.Println("Bot inicializado")

	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	// Configura os comandos do bot
	commands := []tgbotapi.BotCommand{
		{Command: "iniciar", Description: "Inicia uma nova interação"},
		{Command: "cancelar", Description: "Cancela a consulta atual"},
		{Command: "resetar", Description: "Limpa os dados da consulta atual"},
		{Command: "meuid", Description: "Exibe o ID do Telegram"},
		{Command: "ajuda", Description: "Ajuda"},
	}
	cfg := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(cfg); err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.IsCommand() {
				fmt.Println(update.Message.Command())
				println("É um comando")
			}
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID
			// bot.Send(msg)
		}
	}
}
