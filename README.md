# ✈️ telegram-bot-template - Шаблон для телеграм ботов 🤖

<h3 align="left">🛠 Стек технологий:</h3>
<!-- Telegram -->
<a href="https://telegram.org/" target="_blank">
<img src="https://img.icons8.com/color/48/000000/telegram-app--v3.png" alt="telegram" width="40" height="40"/></a>
<!-- Go -->
<a href="https://golang.org" target="_blank"> 
<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go lang" width="40" height="40"/></a>
<!-- Visual Studio Code -->
<a href="https://code.visualstudio.com/" target="_blank">
<img src="https://img.icons8.com/fluent/48/000000/visual-studio-code-2019.png" alt="vs code" width="40" height="40"/></a>
<!-- Docker -->
<a href="https://github.com/hud0shnik/golang-to-do" >
<img src="https://img.icons8.com/fluency/48/000000/docker.png" alt="Docker" width="40" height="40"/></a>

<h3 align="left">📄 О проекте:</h3>
Пустой шаблон для разработки Telegram-бота на Go 

<h3 align="left">🐋 Запуск в Docker:</h3>

Бота можно запустить в Docker-контейнере. Для этого нужно собрать проект:
```
docker build -t tg_bot .
```

И запустить:
```
docker run --name=telegram_bot -e TOKEN="<token_value>" tg_bot
```

