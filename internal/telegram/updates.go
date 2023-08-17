package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// telegramResponse - структура ответа
type telegramResponse struct {
	Result []Update `json:"result"`
}

// Update - структура обновления
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message - структура сообщения
type Message struct {
	Chat    Chat    `json:"chat"`
	Text    string  `json:"text"`
	Sticker Sticker `json:"sticker"`
}

// Sticker - структура стикера
type Sticker struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
}

// Chat - структура чата
type Chat struct {
	ChatId int `json:"id"`
}

// GetUpdates - функция получения апдейтов
func GetUpdates(botUrl string, offset int) ([]Update, error) {

	// Запрос для получения апдейтов
	resp, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Запись и обработка полученных данных
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse telegramResponse
	if err := json.Unmarshal(body, &restResponse); err != nil {
		return nil, err
	}

	return restResponse.Result, nil

}
