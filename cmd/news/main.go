package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"news/internal/spider"
	tg "news/internal/telegram"
)

var params = spider.Setting{}

func init() {
	flag.StringVar(&params.Token, "tk", "", "Telegram Bot API")
	flag.StringVar(&params.Channel, "id", "", "@Name или ID телеграмм группы")
	flag.Int64Var(&params.Timeout, "t", 600, "Timeout Время интервал между запросами на сайт")
}

func main() {
	flag.Parse()

	if params.Token != "" && params.Channel != "" {
		log.Fatalln(app())
	} else {
		log.Println("ошибка: Не указаны флаги")
		os.Exit(1)
	}
}

//app запуск приложения
func app() error {
	xpath := `//div[@data-block-id="1774812525"]//div[@class="cell-list__list"]/div[@class="cell-list__item m-no-image"]/a`
	cache := spider.DataBase{} // Временное хранения новостного поста

	for {
		res := spider.GetData(xpath)

		if res[0].Text != cache.Text {
			cache = spider.DataBase{Text: res[0].Text, URL: res[0].URL}

			message := fmt.Sprintf("%s\n%s\n", res[0].Text, res[0].URL)
			err := tg.Send(params.Token, params.Channel, message)
			if err != nil {
				return errors.New("ошибка: запрос на отправку в Telegram. " + err.Error())
			}
		}

		time.Sleep(time.Duration(params.Timeout) * time.Second)
	}
}
