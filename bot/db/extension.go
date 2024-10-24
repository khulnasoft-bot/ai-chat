package db

import (
	"khulnasoft.com/proxy"
)

func (bot *Bot) GetAvatarURL() string {
	return proxy.BaseURL.JoinPath("bots", bot.ID.String(), "avatar").String()
}
