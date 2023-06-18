package main

import (
	"log"
	"tgBot/internal/handler"
	"tgBot/internal/telegram"

	"github.com/spf13/viper"
)

func main() {

	// Инициализация конфига (токенов)
	err := InitConfig()
	if err != nil {
		log.Println("Config error: ", err)
		return
	}

	// Url бота для отправки и приёма сообщений
	botUrl := "https://api.telegram.org/bot" + viper.GetString("token")
	offSet := 0

	for {

		// Получение апдейтов
		updates, err := telegram.GetUpdates(botUrl, offSet)
		if err != nil {
			log.Println("Something went wrong: ", err)
		}

		// Обработка апдейтов
		for _, update := range updates {
			handler.Respond(botUrl, update)
			offSet = update.UpdateId + 1
		}

		// Вывод апдейтов в консоль для тестов
		// fmt.Println(updates)
	}
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
