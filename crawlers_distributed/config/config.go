package config

const (
	//AppHost     = "localhost"
	AppHost    = "127.0.0.1"
	AppPort    = 1234
	AppAddress = ":1234"
	Worker0    = 9000

	ItemSaverRPCApi = "ItemSavesService.Save"
	CrawlersServiceSaverRPCApi = "CrawlersService.Process"

	ElasticIndex = "dating_profile"

	Qps = 20
)

const (
	FuncParseCityList    = "ParseCityList"
	FuncParseProfileList = "ParseProfileList"
	FuncParseProfile     = "ParseProfile"
	FuncNilParser        = "NilParser"
)
