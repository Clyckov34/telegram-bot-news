package spider

import (
	"github.com/gocolly/colly"
)

//GetData запрос на получения данных. Паук
func GetData(XPath string) []DataBase {
	var spider = NewSpider{
		Visit: "https://ria.ru/",
		Colly: colly.NewCollector(
			colly.AllowedDomains("ria.ru"),
		)}

	spider.newCollector()
	spider.error()
	spider.status()

	return newsMain(spider, XPath)
}

//newsMain парсинг главных новостей
func newsMain(s NewSpider, XPath string) []DataBase {
	var data = make([]DataBase, 0)

	s.Colly.OnXML(XPath, func(x *colly.XMLElement) {
		data = append(data, DataBase{Text: x.Text, URL: x.Attr("href")})
	})

	s.Colly.Visit(s.Visit)
	return data
}
