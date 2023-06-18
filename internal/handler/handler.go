package handler

import (
	"tgBot/internal/commands"
	"tgBot/internal/send"
	"tgBot/internal/telegram"
)

// Обработчик сообщений
func Respond(botUrl string, update telegram.Update) error {

	// msg - текст полученного сообщения
	msg := update.Message.Text
	id := update.Message.Chat.ChatId

	// Обработчик комманд
	switch msg {
	case "/hello":
		send.SendMsg(botUrl, id, "101!")
		return nil
	case "/help":
		commands.Help(botUrl, update)
		return nil
	}

	send.SendMsg(botUrl, id, "Я не понимаю, чтобы узнать список команд, воспользуйтесь /help")
	return nil
}
