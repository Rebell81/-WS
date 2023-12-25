package qBit

import (
	"context"
	"fmt"
	"strings"

	"github.com/autobrr/go-qbittorrent"
)

func GetTorrents(ctx context.Context, client *qbittorrent.Client) (error, []qbittorrent.Torrent) {
	var (
		filter = "seeding"
		//filter   = "all"
		category string
		tag      string
		hashes   []string
	)
	req := qbittorrent.TorrentFilterOptions{
		Filter:   qbittorrent.TorrentFilter(strings.ToLower(filter)),
		Category: category,
		Tag:      tag,
		Hashes:   hashes,
	}

	torrents, err := client.GetTorrentsCtx(ctx, req)
	if err != nil {
		return fmt.Errorf("ERROR: could not get torrents %v", err), nil
	}

	return nil, torrents
}
