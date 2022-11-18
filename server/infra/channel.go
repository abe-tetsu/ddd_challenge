package infra

import (
	"context"
	"database/sql"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	"log"
)

type ChannelRepository struct {
	Conn *sql.DB
}

func NewChannelRepository(conn *sql.DB) *ChannelRepository {
	return &ChannelRepository{Conn: conn}
}

func ScanChannels(rows *sql.Rows) ([]*domain.Channel, int, error) {
	channels := make([]*domain.Channel, 0)

	for rows.Next() {
		var v *domain.Channel
		if err := rows.Scan(&v.ChannelID, &v.CreatedAt, &v.UpdatedAt); err != nil {
			log.Printf("[ERROR] scan ScanChannels: %+v", err)
			return nil, 0, err
		}
		channels = append(channels, v)
	}

	return channels, len(channels), nil
}

func (c ChannelRepository) CreateChannel(ctx context.Context, channel *domain.Channel) error {
	query := "INSERT INTO channel (channel_ID, channel_name, created_at, updated_at) VALUES (?,?,?,?) "
	_, err := c.Conn.ExecContext(ctx, query, channel.ChannelID, channel.CreatedAt, channel.UpdatedAt)
	if err != nil {
		log.Printf("[ERROR] can't create CreateChannel: %+v", err)
		return nil
	}

	return nil
}

func (c ChannelRepository) GetChannels(ctx context.Context) ([]*domain.Channel, error) {
	query := "SELECT * FROM channel"
	rows, err := c.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[ERROR] not found Channels: %+v", err)
		return nil, err
	}

	channels, _, err := ScanChannels(rows)
	if err != nil {
		log.Printf("[ERROR] can't scan Channels: %+v", err)
		return nil, err
	}

	return channels, nil
}

func (c ChannelRepository) UpdateChannel(ctx context.Context, channelID domain.ChannelID, updatedChannel *domain.Channel) error {
	query := "UPDATE channel set ChannelName = ? WHERE ChannelID = ? "
	_, err := c.Conn.ExecContext(ctx, query, updatedChannel, channelID)
	if err != nil {
		log.Printf("[ERROR] can't UpdateChannel: %+v", err)
		return nil
	}

	return nil
}

func (c ChannelRepository) DeleteChannel(ctx context.Context, channelID domain.ChannelID) error {
	query := "DELETE FROM channel WHERE id = ?"
	_, err := c.Conn.ExecContext(ctx, query, channelID)
	if err != nil {
		log.Printf("[ERROR] can't DeleteChannel: %+v", err)
		return nil
	}

	return nil
}
