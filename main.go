package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tgBot/mods"

	"github.com/spf13/viper"
)

// Функция получения апдейтов (сообщений)
func getUpdates(botUrl string, offset int) ([]mods.Update, error) {

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
	var restResponse mods.TelegramResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

// Обработчик сообщений
func respond(botUrl string, update mods.Update) error {

	// msg - текст полученного сообщения
	msg := update.Message.Text

	// Обработчик комманд
	switch msg {
	case "/hello":
		mods.SendMessage(botUrl, update, "101!")
		return nil
	case "/help":
		mods.Help(botUrl, update)
		return nil
	}

	mods.SendMessage(botUrl, update, "Я не понимаю, чтобы узнать список команд, воспользуйтесь /help")
	return nil
}

func main() {

	// Инициализация конфига (токенов)
	err := mods.InitConfig()
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
