package handler

import (
	"strings"
	"tgBot/internal/commands"
	"tgBot/internal/send"
	"tgBot/internal/telegram"
)

// Обработчик сообщений
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
			send.SendMsg(botUrl, chatId, "101!")
		default:
			// Дефолтный ответ
			send.SendMsg(botUrl, chatId, "У меня пока нет такой команды")
		}

	} else {

		// Проверка на стикер
		if update.Message.Sticker.File_id != "" {
			send.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIaImHkPqF8-PQVOwh_Kv1qQxIFpPyfAAJXAAOtZbwUZ0fPMqXZ_GcjBA")
		} else {
			// Если пользователь отправил не сообщение и не стикер:
			send.SendMsg(botUrl, chatId, "Пока я воспринимаю только текст и стикеры")
			send.SendStck(botUrl, chatId, "CAACAgIAAxkBAAIaImHkPqF8-PQVOwh_Kv1qQxIFpPyfAAJXAAOtZbwUZ0fPMqXZ_GcjBA")
		}

	}
}
