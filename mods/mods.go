package mods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// Структуры для работы с Telegram API
type TelegramResponse struct {
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

// Функция для отправки сообщений пользователю
func SendMessage(botUrl string, update Update, msg string) error {
	// Запись того, что и куда отправить
	botMessage := BotMessage{
		ChatId: update.Message.Chat.ChatId,
		Text:   msg,
	}

	// Запись сообщения в json
	buf, err := json.Marshal(botMessage)
	if err != nil {
		fmt.Println("Marshal json error: ", err)
		return err
	}

	// Отправка сообщения
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

// Вывод списка всех команд
func Help(botUrl string, update Update) {
	SendMessage(botUrl, update, "Привет👋🏻, вот список команд:"+
		"\n\n/command - команда 1")
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
