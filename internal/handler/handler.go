package handler

import (
	"strings"
	"tgBot/internal/commands"
	"tgBot/internal/telegram"
)

// Respond - обработчик сообщений
func Respond(botUrl string, update telegram.Update) {

	// Запись айди чата
	chatId := update.Message.Chat.ChatId

	// Обработчик команд
	if update.Message.Text != "" {

		request := strings.Split(update.Message.Text, " ")

		// Вывод реквеста для тестов
		// fmt.Println("request: \t", request)

		switch request[0] {
		case "/start", "/help":
			commands.Help(botUrl, chatId)
		case "/hello":
			telegram.SendMsg(botUrl, chatId, "101!")
		default:
			// Дефолтный ответ
			telegram.SendMsg(botUrl, chatId, "У меня пока нет такой команды")
		}

	} else {

		// Проверка на стикер
		if update.Message.Sticker.FileId != "" {
			telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIaImHkPqF8-PQVOwh_Kv1qQxIFpPyfAAJXAAOtZbwUZ0fPMqXZ_GcjBA")
		} else {
			// Если пользователь отправил не сообщение и не стикер:
			telegram.SendMsg(botUrl, chatId, "Пока я воспринимаю только текст и стикеры")
			telegram.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIaImHkPqF8-PQVOwh_Kv1qQxIFpPyfAAJXAAOtZbwUZ0fPMqXZ_GcjBA")
		}

	}
}
