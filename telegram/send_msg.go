package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMsg(bot *tgbotapi.BotAPI, msg string, chatId int64) error {
	tgMessage := tgbotapi.NewMessage(chatId, msg)
	tgMessage.ParseMode = "markdown"

	_, err := bot.Send(tgMessage)
	if err != nil {
		return err
	}

	return nil
}
