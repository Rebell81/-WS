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
)

func doCheck(ctx context.Context, config *config.Config) error {
	err, torrents := qBit.GetTorrents(ctx, qBittorrentClient)

	//log.Println(fmt.Sprintf("found %d hashes on client...", len(torrents)))

	hashStrings := qBit.GetHashStrings(torrents)
	result, err := rutracker_api.GetIdByHashes(hashStrings, config)
	if err != nil {
		log.Println(err)
		return err
	}

	err = validateHash(ctx, config, result)
	if err != nil {
		return err
	}

	return nil
}

func validateHash(ctx context.Context, config *config.Config, result map[string]*int) error {
	notFoundOnTrackerHashes := make([]string, 0)
	for key, val := range result {
		if val == nil {
			notFoundOnTrackerHashes = append(notFoundOnTrackerHashes, key)
		}
	}

	for _, hashV1 := range notFoundOnTrackerHashes {
		err, props := qBit.GetProperties(ctx, qBittorrentClient, hashV1)
		if err != nil {
			return err
		}

		comment := strings.Replace(props.Comment, ".org", ".net", 1)
		fmt.Println(fmt.Sprintf("%s|%s|%s", props.Name, hashV1, comment))

		if telegramBotInited {
			err = telegram.SendMsg(
				telegramBotClient,
				fmt.Sprintf("Найден мертвый торент: %s|%s|%s", props.Name, hashV1, comment),
				config.ChatId)
			if err != nil {
				log.Printf(err.Error())
			}
		}
	}

	return nil
}
