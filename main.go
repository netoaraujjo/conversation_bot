package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

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

	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Inicia a interação"},
		{Command: "help", Description: "Ajuda"},
	}

	cfg := tgbotapi.NewSetMyCommands(commands...)
	if _, err := bot.Request(cfg); err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.IsCommand() {
				println("É um comando")
			}
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID
			// bot.Send(msg)
		}
	}
}
