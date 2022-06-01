package telegram

import (
	"net/http"
	"net/url"
)

//Send отправка сообщения в телеграмм боту
func Send(token string, chat_id string, message string) (err error) {
	value := url.Values{
		"chat_id": {chat_id},
		"text":    {message},
	}

	_, err = http.PostForm("https://api.telegram.org/bot"+token+"/sendMessage", value)
	if err != nil {
		return err
	}

	return nil
}
