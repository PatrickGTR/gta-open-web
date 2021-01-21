package helper

type MessageData struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func SendMessage(code string, msg string) []MessageData {
	return []MessageData{{code, msg}}
}
