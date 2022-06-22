# Telegram Bot "NEWS RIA.RU"
Telegram Bot + Spider. Выгрузка новостей с интервалом из источника RIA.RU.

### **Описание флагов:**
- `--help` Справочник.
- `--id` ID телеграмм группы.
- `--tk` Token. Telegram BOT API.
- `--t` Интервал Времяни между запросами на сайт. Default: 10 мин. Указывать в минутах.

### **Подготовка Telegram:**
1. Создайте TELEGRAM_BOT с помощью бота: @BotFather
    > Вы получите TELEGRAM_TOKEN_BOT
2. Создайте группу

    2.1. Добавьте в группу пользователя Вашего Telegram Bot 
    
    2.2. Предоставьте Admin права Вашему Telegram Bot.

    2.3. Добавьте в группу бота @getmyid_bot
        
    > Вы получите TELEGRAM_GROUP_ID. Образец: `Current chat ID: -1005602222740` **TELEGRAM_GROUP_ID всегда начинается с МИНУСА**

    2.4. Удалите из группы бота @getmyid_bot
    
    2.5. Отредактируйте права остальным учасникам группы

    2.6. Готово

### **Подготовка утилиты к работе:**
1. [Скачать и установить GO на ПК](https://go.dev/dl/)
2. [Скачать и установить Git на ПК](https://git-scm.com/downloads) 
3. Скачать код c репозитории в свою отдельную папку (С помощью команды в терминале): 

    3.1. Windows > `git clone https://github.com/Clyckov34/telegram-bot-news.git`

    3.2. Linux $ `git clone https://github.com/Clyckov34/telegram-bot-news.git`
4. Build App (Собрать пиложения в бинарный файл)
    
    4.1. Windows > `GOOS=windows GOARCH=amd64 go build -o="main.exe" cmd/news/main.go`

    4.2. Linux $ `GOOS=linux GOARCH=amd64 go build cmd/news/main.go`
5. Run App (Запуск бинарного файла **НЕ ЗАБУДЬТЕ УКАЗАТЬ ФЛАГИ**  Telegram Token, и ID группы)

    5.1. Windows > `.\main.exe --tk=TELEGRAM_TOKEN_BOT --id=TELEGRAM_GROUP_ID`
    
    5.2. Linux $ `УКАЗАТЬ_ПОЛНЫЙ_ПУТЬ_К_ФАЙЛУ/main  --tk=TELEGRAM_TOKEN_BOT --id=TELEGRAM_GROUP_ID`
6. Готово.

### Запуск через DockerFile:
- Build - `docker build -t news .`

- Run -   `docker run news ./app --tk=TELEGRAM_TOKEN_BOT --id=TELEGRAM_GROUP_ID`

### Linux: Автозапуск Telegram Bot systemctl
1. Скомпилируйте приложения ` go build cmd/news/main.go `
2. Разместите скомпилированный файл ` main ` в каталог ` /usr/local/bin/MY_PROJECT`
3. Откройте файл ` telegram-bot-news.service ` Укажите путь к приложению, токен TOKEN_TELEGRAM_BOT, и TELEGRAM_GROUP_ID 
4. Сохраните, и закройте файл
5. Разместите файл ` telegram-bot-news.service ` в каталоге ` /etc/systemd/system/ `
6. Укажите права доступа к файлу ` sudo chmod +x /usr/local/bin/MY_PROJECT/main `
7. Запустите скрипт как службу:
- ` sudo systemctl enable telegram-bot-news.service `

- ` sudo systemctl start telegram-bot-news.service `

### **Примечание !!!**
1. Если в течении минуты не появился новостной пост в Telegram группе, скорее всего изменился адрес блока XPath на сайте https://ria.ru/

    > Решения: Измените на актуальный XPath в файле ` internal/spider/default.go ` в блоке const(xPath). И заново повторите пункт № 4 - 6
    
    > Образец кода XPath: `//div[@data-block-id="1795085849"]//div[@class="cell-list__list"]/div[@class="cell-list__item m-no-image"]/a` 


2. Старайте не делать частых запросов на сайт https://ria.ru/, могут заблокировать на время Ваш IP-адрес за DOS-атаку... 
    > По умолчанию указан интервал в раз 10 мин.