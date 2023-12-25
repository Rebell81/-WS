package qBit

import (
	"context"

	"github.com/autobrr/go-qbittorrent"
)

func GetProps(ctx context.Context, client *qbittorrent.Client, hash string) (error, *qbittorrent.TorrentProperties) {
	props, err := client.GetTorrentProperties(hash)
	if err != nil {
		return err, nil
	}

	return nil, &props
}
