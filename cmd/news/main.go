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
	flag.Int64Var(&params.Timeout, "t", 10, "Timeout (В минутах) Время интервал между запросами на сайт")
}

func main() {
	flag.Parse()

	if len(params.Token) == 0 && len(params.Channel) == 0 {
		log.Fatalln("ошибка: Не указаны флаги")
	}

	if err := spider.App(params); err != nil {
		log.Fatalln(err)
	}
}
