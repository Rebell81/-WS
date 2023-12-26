package qBit

import (
	"context"

	"github.com/autobrr/go-qbittorrent"
)

func GetProperties(ctx context.Context, client *qbittorrent.Client, hash string) (error, *qbittorrent.TorrentProperties) {
	props, err := client.GetTorrentPropertiesCtx(ctx, hash)
	if err != nil {
		return err, nil
	}

	return nil, &props
}
