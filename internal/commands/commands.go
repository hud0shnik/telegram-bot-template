package commands

import (
	"tgBot/internal/send"
	"tgBot/internal/telegram"
)

// Вывод списка всех команд
func Help(botUrl string, update telegram.Update) {
	send.SendMsg(botUrl, update.Message.Chat.ChatId, "Привет👋🏻, вот список команд:"+
		"\n\n/hello - команда 1")
}
