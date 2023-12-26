package sweek

import (
	"context"
	"cws/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/autobrr/go-qbittorrent"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegramBotInited = false
	telegramBotClient *tgbotapi.BotAPI
	qBittorrentClient *qbittorrent.Client
)

func Process(ctx context.Context, config *config.Config) (int, error) {
	err := initClients(ctx, config)
	if err != nil {
		return -1, err
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	for {
		if !config.ManualCheckOnly {
			err := doCheck(ctx, config)
			if err != nil {
				log.Println(err.Error())
			}

			time.Sleep(time.Duration(config.DurationSeconds) * time.Second)
		} else {
			time.Sleep(3600 * time.Second)
		}
	}
}
