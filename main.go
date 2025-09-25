package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	START int64 = iota
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
	State      int64
	ID         int64
	Name       string
	Username   string
	Parameters map[string]string
}

func handleMessageText(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Println("Tratando texto: " + update.Message.Text)
	return 0
}

func handleMessagePhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Println("Tratando foto")
	return 0
}

func handleCommandStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Println("Tratando comando start")
	return 0
}

func handleCommandAjuda(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Println("Tratando comando ajuda")
	return 0
}

func handleCommandCancelar(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Println("Tratando comando cancelar")
	return 0
}

func handleCommandResetar(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Println("Tratando comando resetar")
	return 0
}

func handleCommandMeuID(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	fmt.Printf("Tratando comando meuid. ID do usuário: %d\n", update.Message.From.ID)
	return 0
}

func main() {
	// Bloco de configuração do Bot
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %s", err)
	}

	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Erro ao inicializar o bot: %s", err)
	}

	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
	log.Println("Bot inicializado")
	// fim do bloco de configuração do Bot

	// users := make(map[int64]*Usuario)

	ch := NewConversationHandler()
	ch.EntryPoints = []EventHandler{
		{Match: MessageHandler(TEXT), Handler: handleMessageText},
		{Match: MessageHandler(PHOTO), Handler: handleMessagePhoto},
		{Match: CommandHandler("start"), Handler: handleCommandStart},
		{Match: CommandHandler("ajuda"), Handler: handleCommandAjuda},
		{Match: CommandHandler("cancelar"), Handler: handleCommandCancelar},
		{Match: CommandHandler("resetar"), Handler: handleCommandResetar},
		{Match: CommandHandler("meuid"), Handler: handleCommandMeuID},
	}

	for update := range updates {
		for _, h := range ch.EntryPoints {
			if h.Match(update) {
				h.Handler(bot, update)
			}
		}
	}
}
