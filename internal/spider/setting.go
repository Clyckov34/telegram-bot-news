package spider

//dataBase Структура данных
type dataBase struct {
	Text string
	URL  string
}

//Config параметры запуска приложения
type Config struct {
	Token   string
	Channel string
	Timeout int64
}

//Spider параметры запуска паука
type Spider struct {
	Visit          string
	AllowedDomains string
	XPath          string
}
