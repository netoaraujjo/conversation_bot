package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"github.com/joho/godotenv"
// )

// const (
// 	START int = iota
// 	TIPO_DA_BUSCA
// 	IMEI
// 	CPF
// 	PLACA
// 	CHASSI
// 	MOTOR
// 	MARCA
// 	COR
// )

// func selecionaTipoBusca(bot *tgbotapi.BotAPI, update tgbotapi.Update, usuario *Usuario) {
// 	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
// 	bot.Request(callback)
// 	opcaoSelecionada := update.CallbackData()
// 	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Opção selecionada: %s", opcaoSelecionada))
// 	bot.Send(msg)
// 	usuario.State = START
// }
// func start(bot *tgbotapi.BotAPI, update tgbotapi.Update, usuario *Usuario) {
// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Selecione uma oção abaixo:")
// 	msg.ReplyMarkup = searchOptionsKeyboard
// 	bot.Send(msg)
// 	usuario.State = TIPO_DA_BUSCA
// }

// type Handler struct {
// 	States map[int][]func(bot *tgbotapi.BotAPI, update tgbotapi.Update, usuario *Usuario)
// }

// func newHandler() *Handler {
// 	return &Handler{
// 		States: make(map[int][]func(bot *tgbotapi.BotAPI, update tgbotapi.Update, usuario *Usuario)),
// 	}
// }

// type Usuario struct {
// 	State      int
// 	ID         int
// 	Name       string
// 	Username   string
// 	Parameters map[string]string
// }

// var searchOptionsKeyboard = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Celular", "celular"),
// 		tgbotapi.NewInlineKeyboardButtonData("Pessoa", "pessoa"),
// 	),
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Erro ao carregar variáveis de ambiente: %s", err)
// 	}
// 	botToken := os.Getenv("BOT_TOKEN")

// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Fatalf("Erro ao inicializar o bot: %s", err)
// 	}

// 	log.Println("Bot inicializado")
// 	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

// 	commands := []tgbotapi.BotCommand{
// 		{Command: "iniciar", Description: "Inicia uma nova interação"},
// 		{Command: "cancelar", Description: "Cancela a consulta atual"},
// 		{Command: "resetar", Description: "Limpa os dados da consulta atual"},
// 		{Command: "meuid", Description: "Exibe o ID do Telegram"},
// 		{Command: "ajuda", Description: "Ajuda"},
// 	}
// 	cfg := tgbotapi.NewSetMyCommands(commands...)
// 	if _, err := bot.Request(cfg); err != nil {
// 		log.Panic(err)
// 	}
// 	handler := newHandler()
// 	handler.States[START] = []func(bot *tgbotapi.BotAPI, update tgbotapi.Update, usuario *Usuario){start}
// 	handler.States[TIPO_DA_BUSCA] = selecionaTipoBusca
// 	usuario := &Usuario{
// 		State: 0,
// 	}
// 	for update := range updates {
// 		handler.States[usuario.State](bot, update, usuario)
// 		if update.Message != nil {
// 			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
// 			switch update.Message.Command() {
// 			case "start", "iniciar":
// 			case "meuid":
// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Seu id do Telegram, toque no número para copiar: `%d`", update.Message.Chat.ID))
// 				msg.ReplyMarkup = searchOptionsKeyboard
// 				msg.ParseMode = tgbotapi.ModeMarkdownV2
// 				bot.Send(msg)
// 			}
// 			if update.Message.IsCommand() {
// 				fmt.Println(update.Message.Command())
// 				println("É um comando")
// 			}
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// 			msg.ReplyToMessageID = update.Message.MessageID
// 			bot.Send(msg)
// 		}
// 		if update.CallbackQuery != nil {
// 			fmt.Println(update.CallbackData())
// 			switch update.CallbackData() {
// 			case "pessoa":
// 				fmt.Println("Selecionou a opção Pessoa")
// 			case "celular":
// 				fmt.Println("Selecionou a opção Celular")
// 			}
// 		}
// 	}
// }
