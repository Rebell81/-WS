package sweek

import (
	"context"
	"cws/config"
	"cws/telegram"
	"fmt"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func readIncome(ctx context.Context, config *config.Config) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := telegramBotClient.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Text {
			case "/check":
				err := telegram.SendMsg(telegramBotClient, "Ща поищем, жди", config.ChatId)
				if err != nil {
					log.Println(err)
				}

				err = doCheck(ctx, config)
				if err != nil {
					err = telegram.SendMsg(telegramBotClient, fmt.Sprintf("Ошибка во время ручного обновления: %s", err), config.ChatId)
					if err != nil {
						log.Println(err)
					}
				} else {
					err = telegram.SendMsg(telegramBotClient, "Ручная проверка выполнена", config.ChatId)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}
