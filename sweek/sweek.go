package sweek

import (
	"context"
	"cws/config"
	"cws/qBit"
	"cws/rutracker_api"
	"cws/telegram"
	"fmt"
	"log"
	"strings"

	"github.com/autobrr/go-qbittorrent"
)

func Process(ctx context.Context, config *config.Config) (int, error) {
	bot, err := telegram.InitBot(config.Token)
	var botInited = false

	if err != nil {
		log.Println(fmt.Errorf("failed to init telegram bot: %w", err))
	} else {
		log.Println("telegram bot inited")
		botInited = true
	}

	var qbitCfg = qbittorrent.Config{
		Username: config.Login,
		Password: config.Password,
	}
	if config.SSL {
		qbitCfg.Host = fmt.Sprintf("https://%s:%d", config.Host, config.Port)
	} else {
		qbitCfg.Host = fmt.Sprintf("http://%s:%d", config.Host, config.Port)
	}

	err, qbClient := qBit.InitClient(ctx, qbitCfg)
	if err != nil {
		log.Println(fmt.Errorf("failed to init qBittorrent client: %w", err))

		return -1, err
	}

	log.Println("qBittorrent client created and login success")
	err, torrents := qBit.GetTorrents(ctx, qbClient)

	hashes := make([]string, 0)
	hashStrings := make([]string, 0)
	counter := 0

	for _, v := range torrents {
		hashes = append(hashes, v.InfohashV1)
		if counter == 99 {
			hashStrings = append(hashStrings, strings.Join(hashes, ","))
			hashes = make([]string, 0)
			counter = 0
		} else {
			counter++
		}
	}

	log.Println(fmt.Sprintf("found %d hashes on client...", len(torrents)))

	result, err := rutracker_api.GetIdByHashes(hashStrings, config.Api)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	notFoundOnTrackerHashes := make([]string, 0)
	for key, val := range result {
		if val == nil {
			notFoundOnTrackerHashes = append(notFoundOnTrackerHashes, key)
		}
	}

	for _, hash := range notFoundOnTrackerHashes {
		err, props := qBit.GetProps(ctx, qbClient, hash)
		if err != nil {
			return -1, err
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

	return 0, nil
}
