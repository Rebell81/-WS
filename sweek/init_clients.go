package sweek

import (
	"context"
	"cws/config"
	"cws/telegram"
	"fmt"
	"log"

	"github.com/autobrr/go-qbittorrent"
)

func initClients(ctx context.Context, config *config.Config) (err error) {
	telegramBotClient, err = telegram.InitBot(config.TelegramToken)
	if err != nil {
		log.Println(fmt.Errorf("failed to init telegram bot: %w", err))
	} else {
		log.Println("telegram bot inited")
		telegramBotInited = true
	}

	if telegramBotInited {
		go readIncome(ctx, config)
	}

	var qbitCfg = qbittorrent.Config{
		Username:      config.Login,
		Password:      config.Password,
		TLSSkipVerify: false,
	}
	if config.SSL {
		qbitCfg.Host = fmt.Sprintf("https://%s/", config.Host)
	} else {
		qbitCfg.Host = fmt.Sprintf("http://%s:%d", config.Host, config.Port)
	}

	qBittorrentClient = qbittorrent.NewClient(qbitCfg)

	if err := qBittorrentClient.LoginCtx(ctx); err != nil {
		return fmt.Errorf("failed to init qBittorrent client: %w", err)
	}

	log.Println("qBittorrent client created and login success")

	return nil
}
