package models

type Chat struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	LastMsgTimestamp string `json:"lastMsgTimestamp"`
}
