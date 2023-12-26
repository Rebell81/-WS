package sweek

import (
	"context"
	"cws/config"
	"cws/qBit"
	"cws/rutracker_api"
	"cws/telegram"
	"fmt"
	"github.com/autobrr/go-qbittorrent"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	botInited   = false
	bot         *tgbotapi.BotAPI
	qBittorrent *qbittorrent.Client
)

func Process(ctx context.Context, config *config.Config) (int, error) {
	err := initCLients(ctx, config)
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
			_, err := doCheck(ctx, config)
			if err != nil {
				log.Println(err.Error())
			}

			time.Sleep(60 * time.Second)
		} else {
			time.Sleep(3600 * time.Second)
		}
	}
}

func initCLients(ctx context.Context, config *config.Config) (err error) {
	bot, err = telegram.InitBot(config.TelegramToken)
	if err != nil {
		log.Println(fmt.Errorf("failed to init telegram bot: %w", err))
	} else {
		log.Println("telegram bot inited")
		botInited = true
	}

	go readIncome(ctx, config)

	var qbitCfg = qbittorrent.Config{
		Username: config.Login,
		Password: config.Password,
	}
	if config.SSL {
		qbitCfg.Host = fmt.Sprintf("https://%s:%d", config.Host, config.Port)
	} else {
		qbitCfg.Host = fmt.Sprintf("http://%s:%d", config.Host, config.Port)
	}

	qBittorrent = qbittorrent.NewClient(qbitCfg)

	if err := qBittorrent.LoginCtx(ctx); err != nil {
		return fmt.Errorf("failed to init qBittorrent client: %w", err)
	}

	log.Println("qBittorrent client created and login success")

	return nil
}

func readIncome(ctx context.Context, config *config.Config) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Text {
			case "/check":
				err := telegram.SendMsg(bot, "Ща поищем, жди", config.ChatId)
				if err != nil {
					log.Println(err)
				}

				res, err := doCheck(ctx, config)
				if err != nil {
					err = telegram.SendMsg(bot, fmt.Sprintf("Ошибка во время ручного обновления: %s", err), config.ChatId)
					if err != nil {
						log.Println(err)
					}
				} else {
					if res {
						err = telegram.SendMsg(bot, "Мертвых торентов не найдено", config.ChatId)
						if err != nil {
							log.Println(err)
						}
					}
				}
			}
		}
	}
}

func doCheck(ctx context.Context, config *config.Config) (bool, error) {
	err, torrents := qBit.GetTorrents(ctx, qBittorrent)

	log.Println(fmt.Sprintf("found %d hashes on client...", len(torrents)))

	hashStrings := qBit.GetHashesStrings(torrents)
	result, err := rutracker_api.GetIdByHashes(hashStrings, config.RutrackerApiToken)
	if err != nil {
		log.Println(err)
		return false, err
	}

	notFoundOnTrackerHashes := make([]string, 0)
	for key, val := range result {
		if val == nil {
			notFoundOnTrackerHashes = append(notFoundOnTrackerHashes, key)
		}
	}

	for _, hash := range notFoundOnTrackerHashes {
		err, props := qBit.GetProps(ctx, qBittorrent, hash)
		if err != nil {
			return false, err
		}

		com := strings.Replace(props.Comment, ".org", ".net", 1)
		fmt.Println(fmt.Sprintf("%s|%s|%s", props.Name, hash, com))

		if botInited {
			err = telegram.SendMsg(bot, fmt.Sprintf("Найден мертвый торент: %s|%s|%s", props.Name, hash, com), config.ChatId)
			if err != nil {
				log.Printf(err.Error())
			}
		}
	}

	if len(notFoundOnTrackerHashes) == 0 {
		log.Println("all torrents are up to date")
	}

	return true, nil
}
