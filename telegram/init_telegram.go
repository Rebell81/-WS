package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func InitBot(token string) (*tgbotapi.BotAPI, error) {
	fmt.Println(fmt.Sprintf("try init bot '%s'", token))
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}
