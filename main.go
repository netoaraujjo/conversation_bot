package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	START int = iota
	TIPO_DA_BUSCA
	IMEI
	CPF
	PLACA
	CHASSI
	MOTOR
	MARCA
	COR
)

type Usuario struct {
	State      int
	ID         int
	Name       string
	Username   string
	Parameters map[string]string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %s", err)
	}
	botToken := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Erro ao inicializar o bot: %s", err)
	}
	log.Println("Bot inicializado")

	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updates {
	}
}
