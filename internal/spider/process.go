package spider

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gocolly/colly"
)

//App запуск приложения
func App(p Config) error {
	spider := new(Spider)
	spider.Visit = visit
	spider.AllowedDomains = allowedDomains
	spider.XPath = xPath

	var cacheNews = dataBase{} // Временное хранения новостного поста

	for {
		res := getData(spider)[0]

		if res.Text != cacheNews.Text {
			cacheNews = dataBase{Text: res.Text, URL: res.URL}

			message := fmt.Sprintf("%s\n%s\n", res.Text, res.URL)
			if err := send(p, message); err != nil {
				return err
			}
		}
		time.Sleep(time.Duration(p.Timeout) * time.Minute)
	}
}

//getData запрос на получения данных. Паук
func getData(s *Spider) []dataBase {
	var data = make([]dataBase, 0)

	c := colly.NewCollector(colly.AllowedDomains(s.AllowedDomains))

	c.OnXML(s.XPath, func(x *colly.XMLElement) {
		data = append(data, dataBase{Text: x.Text, URL: x.Attr("href")})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Запрос OK: -", r.Request.URL)
	})

	c.Visit(s.Visit)

	return data
}

//send отправка сообщения в телеграмм боту
func send(p Config, message string) (err error) {
	value := url.Values{
		"chat_id": {p.Channel},
		"text":    {message},
	}

	_, err = http.PostForm("https://api.telegram.org/bot"+p.Token+"/sendMessage", value)
	if err != nil {
		return err
	}

	return nil
}
