package unimes

import "time"

// Unimes
type UniversalMessage struct {
	ID        int64       `json:"id"`        //message ID
	Time      time.Time   `json:"time"`      //Дата получения
	IsPublic  bool        `json:"is_public"` //Указывает доступность сообщения сторонним пользователям
	Sender    Destination `json:"sender"`    //Создатель сообщения
	Location  Destination `json:"location"`  //Место где было получено сообщение
	Recipient Destination `json:"recipient"` //Кем получено сообщение
	Content   Content     `json:"content"`   //Содержание сообщения
}

type Content struct {
	Type        TypeContent `json:"type"`
	MessageLink string      `json:"msg_link"`
	FotoLink    string      `json:"foto_link"`
	VideoLink   string      `json:"video_link"`
	Text        string      `json:"text"`
}

type Destination struct {
	ID        int64       `json:"id"`
	Type      string      `json:"type"`    //тип адрессата(пользователь, канал, группа)
	Service   TypeService `json:"service"` //телеграм, фейсбук, файловая система
	FirstName string      `json:"first_name"`
	Lastname  string      `json:"last_name"`
	Username  string      `json:"username"`
	Phone     string      `json:"phone"`
}

type TypeService string

const (
	TypeServiceTelegram TypeService = "tg"
)

type TypeContent string

const (
	TypeContentText  TypeContent = "text"
	TypeContentPhoto TypeContent = "photo"
	TypeContentVideo TypeContent = "video"
)

/*
//тип адрессата(пользователь, канал, группа)
type TypeDestination string

const (
	TypeDestinationChannel TypeDestination = "channel"
	TypeDestinationUser    TypeDestination = "user"
	TypeDestinationGroup   TypeDestination = "group"
	TypeDestinationBot     TypeDestination = "bot"
	TypeDestinationUserbot TypeDestination = "userbot"
)
*/

func New() *UniversalMessage {
	return &UniversalMessage{}
}
