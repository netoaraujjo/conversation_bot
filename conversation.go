package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Usuario struct {
	State      int64
	ID         int64
	Name       string
	Username   string
	Parameters map[string]string
}

func handleMessageText(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Println("Tratando texto: " + update.Message.Text)
	return 0
}

func handleMessagePhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Println("Tratando foto")
	return 0
}

func handleCommandStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Println("Tratando comando start")
	fmt.Println("Selecione o tipo de busca")

	searchOptionsKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Celular", "celular"),
			tgbotapi.NewInlineKeyboardButtonData("Pessoa", "pessoa"),
		),
	)

	msg := tgbotapi.NewMessage(update.FromChat().ID, "Selecione o tipo de busca:")
	msg.ReplyMarkup = searchOptionsKeyboard
	bot.Send(msg)

	return TIPO_DA_BUSCA
}

func handleCommandAjuda(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Println("Tratando comando ajuda")
	return 0
}

func handleCommandCancelar(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Println("Tratando comando cancelar")
	return 0
}

func handleCommandResetar(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Println("Tratando comando resetar")
	return 0
}

func handleCommandMeuID(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	fmt.Printf("Tratando comando meuid. ID do usuário: %d\n", update.Message.From.ID)
	return 0
}

func handleTipoBusca(bot *tgbotapi.BotAPI, update tgbotapi.Update) int64 {
	bot.Request(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	tipoBusca := update.CallbackData()
	fmt.Printf("Opção selecionada: %s\n", tipoBusca)
	return END
}
