package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
func sendMessage(botUrl string, update Update, msg string) error {
	// Запись того, что и куда отправить
	botMessage := BotMessage{
		ChatId: update.Message.Chat.ChatId,
		Text:   msg,
	}

	// Запись сообщения в json
	buf, err := json.Marshal(botMessage)
	if err != nil {
		log.Println("Marshal json error: ", err)
		return err
	}

	// Отправка сообщения
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

// Вывод списка команд
func help(botUrl string, update Update) {
	sendMessage(botUrl, update, "Привет👋🏻, вот список команд:"+
		"\n\n/command - команда 1")
}

func main() {
	log.Println("Config error: ")
	// Инициализация конфига (токенов)
	err := initConfig()
	if err != nil {
		log.Println("Config error: ", err)
		return
	}
	// Url бота для отправки и приёма сообщений
	botUrl := "https://api.telegram.org/bot" + viper.GetString("token")
	offSet := 0

	for {
		// Получение апдейтов
		updates, err := getUpdates(botUrl, offSet)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}

		// Обработка апдейтов
		for _, update := range updates {
			respond(botUrl, update)
			offSet = update.UpdateId + 1
		}

		// Вывод апдейтов в консоль для тестов
		// fmt.Println(updates)
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	// Rest запрос для получения апдейтов
	resp, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Запись и обработка полученных данных
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse TelegramResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

// Обработчик сообщений
func respond(botUrl string, update Update) error {
	// msg - текст полученного сообщения
	msg := update.Message.Text

	// Обработчик комманд
	switch msg {
	case "/command":
		sendMessage(botUrl, update, "101!")
		return nil
	case "/help":
		help(botUrl, update)
		return nil
	}

	sendMessage(botUrl, update, "Я не понимаю, чтобы узнать список команд, воспользуйтесь /help")
	return nil
}

// Функция инициализации конфига (всех токенов)
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
