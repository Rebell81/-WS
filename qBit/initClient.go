package qBit

import (
	"context"
	"fmt"

	"github.com/autobrr/go-qbittorrent"
)

func InitClient(ctx context.Context, cfg qbittorrent.Config) (error, *qbittorrent.Client) {
	qbClient := qbittorrent.NewClient(cfg)

	if err := qbClient.LoginCtx(ctx); err != nil {
		return fmt.Errorf("connection failed: %w|%v", err, cfg), nil
	}

	return nil, qbClient
}
