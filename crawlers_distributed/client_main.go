package main

import (
	"./persist/client"
	"flag"
	"github.com/gpmgo/gopm/modules/log"
	"net/rpc"
	"strings"

	"../crawler/engine"
	"../crawler/model"
	"../crawler/parser/zhenai"
	"../crawler/persist"
	"../crawler/scheduler"
	"./rpcsupport"
	client2 "./worker/client"
)

const url = "http://www.zhenai.com/zhenghun"
const shUrl = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	run()
	//saveToElasticsearch()
	//inseartTestData()
}

var index = "dating_profile"

var (
	itemSaverHost = flag.String("itemSaver_host", "", "ItemSaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker host(,split)")
)
//运行脚本：go run main.go --itemSaver_host=":" --worker_host=":9000,:9001,:9002,:9003"
func run() {
	flag.Parse()
	//itemChan, err := client.ItemServer(config.AppAddress)
	itemChan, err := client.ItemServer(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	processor := client2.CreateProcessor(CreateProcessorPool(strings.Split(*workerHosts, ",")))

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    url,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList")})
}

func CreateProcessorPool(hosts [] string) chan *rpc.Client {
	var clients [] *rpc.Client
	for _, host := range hosts {
		rpcClient, err := rpcsupport.NewRpcClient(host)
		if err == nil {
			clients = append(clients, rpcClient)
			log.Print(0, "Connected to %s", host)
		} else {
			log.Error("Error connecting to %s \t %v", host, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}

func inseartTestData() {
	client, _ := persist.NewElasticClient()
	persist.Save(index, client, engine.Item{
		Id:   "test_id",
		Url:  "http://www.baidu.com",
		Type: "zhenai",
		Payload: model.Profile{
			Age:        34,
			Height:     166,
			Weight:     61,
			Income:     "2000-5000",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "水平",
			Occupation: "人事/行政",
			Marriage:   "未婚",
			House:      "已购房",
			HuKou:      "上海",
			Education:  "大学本科",
			/*Car:        "已购车",*/
		},
	})
}

/*func saveToElasticsearch() {
	url := ""
	profile := model.Profile{
		Age:        34,
		Height:     166,
		Weight:     61,
		Income:     "2000-5000",
		Gender:     "女",
		Name:       "安静的雪",
		Xinzuo:     "水平",
		Occupation: "人事/行政",
		Marriage:   "未婚",
		House:      "已购房",
		HuKou:      "上海",
		Education:  "大学本科",
		Car:        "已购车",
	}
	item := engine.Item{
		Url:     url,
		_Id:      parser.ExtractId(url),
		Payload: profile,
	}

	err := persist.Save(item)
	fmt.Println(id, err)
}*/
