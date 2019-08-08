package persist

import (
	"fmt"
	"strings"
	"testing"
)

func TestItemServer(t *testing.T) {
	/*profile := model.Profile{
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
	Save(profile)*/
	url := `http://album.zhenai.com/u/fsdafsdfsdfk`
	var splits = strings.Split(url, "/")
	fmt.Println(splits[len(splits)-1])
}
