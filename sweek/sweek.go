package sweek

import (
	"context"
	"cws/config"
	"cws/qBit"
	"cws/rutracker_api"
	"fmt"
	"github.com/autobrr/go-qbittorrent"
	"log"
	"strings"
)

func Process(ctx context.Context, config *config.Config) (int, error) {
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
		log.Println(err)

		return -1, err
	}

	log.Println("client created and login success")
	err, torrents := qBit.GetTorrents(ctx, qbClient)

	hashes := make([]string, 0)
	hashStrings := make([]string, 0)
	counter := 0

	for _, v := range torrents {
		hashes = append(hashes, v.Hash)
		if counter == 99 {
			hashStrings = append(hashStrings, strings.Join(hashes, ","))
			hashes = make([]string, 0)
			counter = 0
		} else {
			counter++
		}
	}
	log.Println(fmt.Sprintf("found %d hashes", len(torrents)))

	result, err := rutracker_api.GetIdByHashes(hashStrings, config.Api)

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
		fmt.Println(fmt.Sprintf("%s|%s", hash, com))
	}

	return 0, nil
}
