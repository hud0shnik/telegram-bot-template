package main

import (
	"os"
	"tgBot/internal/handler"
	"tgBot/internal/telegram"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	// Настройка логгера
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	// Загрузка переменных окружения
	godotenv.Load()

	// Url бота для отправки и приёма сообщений
	botUrl := "https://api.telegram.org/bot" + os.Getenv("TOKEN")
	offSet := 0

	// Уведомление о старте
	logrus.Info("Bot is running")

	for {

		// Получение апдейтов
		updates, err := telegram.GetUpdates(botUrl, offSet)
		if err != nil {
			logrus.Println("Something went wrong: ", err)
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
