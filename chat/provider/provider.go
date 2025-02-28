package provider

import (
	"time"

	"khulnasoft.com/bot/db"
	db2 "khulnasoft.com/chat/service/db"
	"encore.dev/types/uuid"
)

type ListChannelsResponse struct {
	Channels []ChannelInfo
}

type SendMessageRequest struct {
	ID      string
	Content string
	Bot     *db.Bot
	UserID  string
	Type    string
}

type ListMessagesResponse struct {
	Messages []*Message
}

type ListMessagesRequest struct {
	FromTimestamp string
	FromMessageID string
}

type UserID = string

type ChannelID = string

// Message is a message sent by a user in a provider channel
type Message struct {
	Provider   db2.Provider
	ProviderID string
	ChannelID  ChannelID
	Author     User
	Content    string
	Time       time.Time
	Type       string
	Bots       []uuid.UUID
}

// User is a user in a provider
type User struct {
	ID      UserID
	Name    string
	Profile string
	BotID   uuid.UUID
}

// ChannelInfo is information about a channel in a provider
type ChannelInfo struct {
	Provider db2.Provider
	ID       ChannelID
	Name     string
}
