package spider

import (
	"log"

	"github.com/gocolly/colly"
)

type DataBase struct {
	Text string
	URL  string
}

type Setting struct {
	Token   string
	Channel string
	Timeout int64
}

type NewSpider struct {
	Visit string
	Colly *colly.Collector
}

//newCollector создается новый паук
func (m *NewSpider) newCollector() {
	m.Colly = colly.NewCollector()
}

//error обработчик ошибок
func (m *NewSpider) error() {
	m.Colly.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
}

//status статистика в реальном времяни
func (m *NewSpider) status() {
	m.Colly.OnScraped(func(r *colly.Response) {
		log.Println("Запрос OK: -", r.Request.URL)
	})
}
