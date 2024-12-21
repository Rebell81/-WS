package qBit

import (
	"context"
	"fmt"

	"github.com/autobrr/go-qbittorrent"
)

func GetTrackers(ctx context.Context, client *qbittorrent.Client, hash string) ([]qbittorrent.TorrentTracker, error) {
	trackers, err := client.GetTorrentTrackersCtx(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("ERROR: could not get trackers %v", err)
	}

	return trackers, nil
}
