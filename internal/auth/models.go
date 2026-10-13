package models

import "time"

type Message struct {
	Text      string    `json:"text"`
	Channel   string    `json:"channel"`
	Timestamp time.Time `json:"ts"`
	ThreadTS  string    `json:"thread_ts,omitempty"`
}

func NewMessage(channelID, text string) *Message {
	return &Message{
		Channel:   channelID,
		Text:      text,
		Timestamp: time.Now(),
	}
}
