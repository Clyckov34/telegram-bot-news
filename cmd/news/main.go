package main

import (
	"flag"
	"log"

	"news/internal/spider"
)

var params = spider.Config{}

func init() {
	flag.StringVar(&params.Token, "tk", "", "Telegram Bot API")
	flag.StringVar(&params.Channel, "id", "", "@Name или ID телеграмм группы")
	flag.Int64Var(&params.Timeout, "t", 600, "Timeout Время интервал между запросами на сайт")
}

func main() {
	flag.Parse()

	if len(params.Token) == 0 && len(params.Channel) == 0 {
		log.Println("ошибка: Не указаны флаги")
	}

	log.Fatalln(spider.App(params))
}
