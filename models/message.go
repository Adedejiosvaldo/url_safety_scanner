package models

type Message struct {
	ChannelID string    `json:"channel_id"`
	Settings  []Setting `json:"settings"`
	Message   string    `json:"message"`
}

type Setting struct {
	Label   string      `json:"label"`
	Type    string      `json:"type"`
	Default interface{} `json:"default"`
}
