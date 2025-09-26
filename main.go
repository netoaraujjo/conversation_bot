package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const (
	END   int64 = -1
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

func NewConversation() *ConversationHandler {
	ch := NewConversationHandler()
	ch.EntryPoints = []EventHandler{
		{Match: MessageHandler(TEXT), Handler: handleMessageText},
		{Match: MessageHandler(PHOTO), Handler: handleMessagePhoto},
		{Match: CommandHandler("start"), Handler: handleCommandStart},
		{Match: CommandHandler("iniciar"), Handler: handleCommandStart},
		{Match: CommandHandler("ajuda"), Handler: handleCommandAjuda},
		{Match: CommandHandler("cancelar"), Handler: handleCommandCancelar},
		{Match: CommandHandler("resetar"), Handler: handleCommandResetar},
		{Match: CommandHandler("meuid"), Handler: handleCommandMeuID},
	}
	ch.States[TIPO_DA_BUSCA] = []EventHandler{
		{Match: CallbackQueryHandler("pessoa"), Handler: handleTipoBusca},
		{Match: CallbackQueryHandler("celular"), Handler: handleTipoBusca},
	}
	return ch
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

	convHandler := NewConversation()

	for update := range updates {
		convHandler.HandleUpdate(bot, update)
	}
}
