package estype

import (
	"encoding/json"
	"fmt"
)

// PornHub Post
type PornhubPost struct {
	Title      string   `json:"title"`
	URL        string   `json:"url"`
	Thumb      string   `json:"thumb"`
	Rate       string   `json:"rate"`
	Duration   string   `json:"duration"`
	Categories []string `json:"categories"`
	Tags       []string `json:"tags"`
}

type ChatTelegram struct {
	ID          int64    `json:"id"`
	Address     string   `json:"address"`
	Name        string   `json:"name"`
	Rating      int      `json:"rating"`
	Observable  bool     `json:"observable"`
	ObserverID  string   `json:"observer_id"`
	Description string   `json:"description"`
	MemberCount int32    `json:"member_count"`
	Lang        string   `json:"lang"`
	Timestamp   string   `json:"timestamp"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
}

// Chat for Connect
type ChatConnector struct {
	ID       string   `json:"id"`
	Address  string   `json:"address"`
	Observer string   `json:"observer"`
	Priority int      `json:"priority"`
	Status   string   `json:"status"`
	Lang     string   `json:"lang"`
	Tags     []string `json:"tags"`
}

/*
type ChatTelegramType string
const (
	ChatTelegramTypeGroup   ChatTelegramType = "group"
	ChatTelegramTypeChannel ChatTelegramType = "channel"
	ChatTelegramTypePrivate ChatTelegramType = "private"
)
*/

func ChatConnectorFromInterfaces(source map[string]interface{}) (*ChatConnector, error) {
	var chat ChatConnector
	data, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = json.Unmarshal(data, &chat)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return &chat, nil
}

func StructToInterfaces(chat interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(chat)
	if err != nil {
		panic(err)
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
