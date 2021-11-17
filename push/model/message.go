package model

type MessageRequest struct {
	ValidateOnly bool     `json:"validate_only"`
	Message      *Message `json:"message"`
}

type MessageResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

type Message struct {
	Data         string         `json:"data,omitempty"`
	Notification *Notification  `json:"notification,omitempty"`
	Android      *AndroidConfig `json:"android,omitempty"`
	Token        []string       `json:"token,omitempty"`
}

type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Image string `json:"image,omitempty"`
}

func NewTransparentMsgRequest() *MessageRequest {
	msgRequest := getDefaultMsgRequest()
	return msgRequest
}

func NewNotificationMsgRequest() *MessageRequest {
	msgRequest := getDefaultMsgRequest()
	msgRequest.Message.Notification = getDefaultNotification()
	return msgRequest
}

func getDefaultMsgRequest() *MessageRequest {
	return &MessageRequest{
		ValidateOnly: false,
		Message: &Message{
			Data: "This is a transparent message data",
		},
	}
}

func getDefaultNotification() *Notification {
	return &Notification{
		Title: "notification tittle",
		Body:  "This is a notification message body",
	}
}
