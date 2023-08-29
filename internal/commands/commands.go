package commands

import "tgBot/internal/telegram"

// Help - вывод списка всех команд
func Help(botUrl string, ChatId int) {
	telegram.SendMsg(botUrl, ChatId, "Привет👋🏻, вот список команд:\n\n"+
		"/help - список команд\n"+
		"/hello - команда 1\n")
}
