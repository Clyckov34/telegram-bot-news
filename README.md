# Telegram Bot "NEWS RIA.RU"
Telegram Bot - Новости из источника RIA.RU

### Запуск утилиты:
- `go run cmd/news/main.go --tk=TOKEN_TELEGRAM_BOT --id=ID_TELEGRAM_GROUP` 

### Запуск утилиты c подробными флагами:
- `go run cmd/telegramBot/main.go --help`

### DockerFile:
- Build - `docker build -t news .`

- Run -   `docker run news ./app --tk=TOKEN_TELEGRAM_BOT --id=ID_TELEGRAM_GROUP`

### Linux: Автозапуск Telegram Bot systemctl
1. Скомпилируйте приложения ` go build cmd/news/main.go `
2. Разместите скомпилированный файл ` main ` в каталог ` /usr/local/bin/MY_PROJECT`
3. Откройте файл ` telegram-bot-news.service ` Укажите путь к приложению, и токен Telegram Bot
4. Сохраните, и закройте файл
5. Разместите файл ` telegram-bot-news.service ` в каталоге /etc/systemd/system/
6. Запустите скрипт как службу:
- ` sudo systemctl enable telegram-bot-news.service `

- ` sudo systemctl start telegram-bot-news.service `
