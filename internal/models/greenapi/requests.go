package greenapi

type SendMessageRequest struct {
	ChatIdOrNumber string `json:"chatId"`
	Message        string `json:"message"`
}

type SendFileRequest struct {
	ChatIdOrNumber string `json:"chatId"`
	UrlFile        string `json:"urlFile"`
	FileName       string `json:"fileName"`
	Caption        string `json:"caption"`
}
