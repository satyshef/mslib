package unimes

// Unimes
type UniversalMessage struct {
	ID       int64    `json:"id"` //message ID
	Metadata Metadata `json:"metadata"`
	//Time      time.Time   `json:"time"`      //Дата получения
	Time      string      `json:"time"`      //Дата получения
	IsPublic  bool        `json:"is_public"` //Указывает доступность сообщения сторонним пользователям
	Sender    Destination `json:"sender"`    //Создатель сообщения
	Location  Destination `json:"location"`  //Место где было получено сообщение
	Recipient Destination `json:"recipient"` //Кем получено сообщение
	Content   Content     `json:"content"`   //Содержание сообщения
}

type Metadata struct {
	Status            StatusMessage   `json:"status"`        //Тип сообщения новое, пересланое, ответ на другое сообщение
	OriginalLink      string          `json:"original_link"` //ссылка на оригинал сообщения
	OriginalChatID    int64           `json:"original_chat_id"`
	OriginalMessageID int64           `json:"original_message_id"`
	OriginalChatType  TypeDestination `json:"original_chat_type"`
	OriginalChatName  string          `json:"original_chat_name"`
}

type Content struct {
	Type      TypeContent `json:"type"`
	Link      string      `json:"link"`
	FotoLink  string      `json:"foto_link"`
	VideoLink string      `json:"video_link"`
	Text      string      `json:"text"`
	Format    string      `json:"format"`
}

type Destination struct {
	ID        int64           `json:"id"`
	Type      TypeDestination `json:"type"`    //тип адрессата(пользователь, канал, группа)
	Service   TypeService     `json:"service"` //телеграм, фейсбук, файловая система
	FirstName string          `json:"first_name"`
	Lastname  string          `json:"last_name"`
	Username  string          `json:"username"`
	Phone     string          `json:"phone"`
}

type StatusMessage string

const (
	StatusMessageNew     StatusMessage = "new"
	StatusMessageReply   StatusMessage = "reply"
	StatusMessageForward StatusMessage = "forward"
)

type TypeService string

const (
	TypeServiceTelegram TypeService = "tg"
)

type TypeContent string

const (
	TypeContentText      TypeContent = "text"
	TypeContentPhoto     TypeContent = "photo"
	TypeContentVideo     TypeContent = "video"
	TypeContentVideoNote TypeContent = "videonote"
)

type TypeDestination string

const (
	TypeDestinationChannel        TypeDestination = "channel"
	TypeDestinationPrivateChannel TypeDestination = "private_channel"
	TypeDestinationGroup          TypeDestination = "group"
	TypeDestinationPrivateGroup   TypeDestination = "private_group"
	TypeDestinationUser           TypeDestination = "user"
	TypeDestinationUserBot        TypeDestination = "userbot"
	TypeDestinationBot            TypeDestination = "bot"
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
