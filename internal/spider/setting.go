package spider

//dataBase Структура данных
type dataBase struct {
	Text string
	URL  string
}

type Config struct {
	Token   string
	Channel string
	Timeout int64
	Spider  Spider
}

type Spider struct {
	Visit          string
	AllowedDomains string
	XPath          string
}
