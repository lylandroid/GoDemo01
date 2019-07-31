package parser

import (
	"../../fetcher"
	"fmt"
	"regexp"
	"testing"
)

const url = "http://www.zhenai.com/zhenghun"
const cityRe2 = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//const rootNodeProfileRe2 = `<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`

const profileRe2 = `<div class="CONTAINER" style="width:100%;margin:20px auto 0;"><div class="m-userInfoDetail" data-v-bff6f798=""><div class="m-title" data-v-bff6f798="">内心独白</div> <div class="m-content-box m-des" data-v-bff6f798=""><span data-v-bff6f798="">彼此理解，支持，尊重，珍惜，更能懂得相互包容！！</span><!----></div> <div class="m-title" data-v-bff6f798="">个人资料</div> <div class="m-content-box" data-v-bff6f798=""><div class="purple-btns" data-v-bff6f798=""><div class="m-btn purple" data-v-bff6f798="">离异</div><div class="m-btn purple" data-v-bff6f798="">39岁</div><div class="m-btn purple" data-v-bff6f798="">魔羯座(12.22-01.19)</div><div class="m-btn purple" data-v-bff6f798="">160cm</div><div class="m-btn purple" data-v-bff6f798="">50kg</div><div class="m-btn purple" data-v-bff6f798="">工作地:阿坝茂县</div><div class="m-btn purple" data-v-bff6f798="">月收入:5-8千</div><div class="m-btn purple" data-v-bff6f798="">零售店店长</div><div class="m-btn purple" data-v-bff6f798="">中专</div></div> <div class="pink-btns" data-v-bff6f798=""><div class="m-btn pink" data-v-bff6f798="">汉族</div><div class="m-btn pink" data-v-bff6f798="">籍贯:四川南充</div><div class="m-btn pink" data-v-bff6f798="">体型:一般</div><div class="m-btn pink" data-v-bff6f798="">不吸烟</div><div class="m-btn pink" data-v-bff6f798="">社交场合会喝酒</div><div class="m-btn pink" data-v-bff6f798="">和家人同住</div><div class="m-btn pink" data-v-bff6f798="">未买车</div><div class="m-btn pink" data-v-bff6f798="">有孩子且住在一起</div><div class="m-btn pink" data-v-bff6f798="">是否想要孩子:视情况而定</div><div class="m-btn pink" data-v-bff6f798="">何时结婚:时机成熟就结婚</div></div></div> <div class="m-title" data-v-bff6f798="">兴趣爱好</div> <div class="m-content-box m-interest f-cl" data-v-bff6f798=""><div class="item f-fl" data-v-bff6f798=""><div class="question f-fl" data-v-bff6f798="">欣赏的一个名人：</div> <div class="answer f-fl" data-v-bff6f798="">马云</div></div><div class="item f-fl" data-v-bff6f798=""><div class="question f-fl" data-v-bff6f798="">喜欢的一首歌：</div> <div class="answer f-fl" data-v-bff6f798="">站着等你三千年</div></div><div class="item f-fl" data-v-bff6f798=""><div class="question f-fl" data-v-bff6f798="">喜欢的一道菜：</div> <div class="answer f-fl" data-v-bff6f798="">未填写</div></div><div class="item f-fl" data-v-bff6f798=""><div class="question f-fl" data-v-bff6f798="">喜欢的一本书：</div> <div class="answer f-fl" data-v-bff6f798="">未填写</div></div><div class="item f-fl" data-v-bff6f798=""><div class="question f-fl" data-v-bff6f798="">喜欢做的事：</div> <div class="answer f-fl" data-v-bff6f798="">未填写</div></div></div> <div class="m-title" data-v-bff6f798="">择偶条件</div> <div class="m-content-box" data-v-bff6f798=""><div class="gray-btns" data-v-bff6f798=""><div class="m-btn" data-v-bff6f798="">30-40岁</div><div class="m-btn" data-v-bff6f798="">155cm以上</div><div class="m-btn" data-v-bff6f798="">工作地:四川阿坝茂县</div><div class="m-btn" data-v-bff6f798="">中专</div><div class="m-btn" data-v-bff6f798="">体型:一般</div><div class="m-btn" data-v-bff6f798="">可以喝酒</div><div class="m-btn" data-v-bff6f798="">不要吸烟</div></div></div> <div class="m-title" data-v-bff6f798="">他的动态</div> <div class="m-content-box m-news" data-v-bff6f798=""><p data-v-bff6f798="">该用户还发布了<span data-v-bff6f798="">9条</span>动态分享他的生活<br data-v-bff6f798="">扫描下载珍爱APP查看他的动态</p> <div class="app" data-v-bff6f798=""></div></div></div></div>`
const rootNodeProfileRe2 = `<div class="m-btn purple" [^>]*>([^<]+)</div>`

var compile2 = regexp.MustCompile(rootNodeProfileRe2)

func TestParseCityList(t *testing.T) {
	//engine.Run(engine.Request{Url:url,engine.ParseResult{}})
	/*compile := regexp.MustCompile(rootNodeProfileRe2)
	str := `<div class="m-content-box" data-v-bff6f798=""><div class="purple-btns" data-v-bff6f798=""><div class="m-btn purple" data-v-bff6f798="">离异</div><div class="m-btn purple" data-v-bff6f798="">30岁</div><div class="m-btn purple" data-v-bff6f798="">天秤座(09.23-10.22)</div><div class="m-btn purple" data-v-bff6f798="">158cm</div><div class="m-btn purple" data-v-bff6f798="">45kg</div><div class="m-btn purple" data-v-bff6f798="">工作地:阿坝汶川</div><div class="m-btn purple" data-v-bff6f798="">月收入:3-5千</div><div class="m-btn purple" data-v-bff6f798="">自由职业</div><div class="m-btn purple" data-v-bff6f798="">大专</div></div> <div class="pink-btns" data-v-bff6f798=""><div class="m-btn pink" data-v-bff6f798="">藏族</div><div class="m-btn pink" data-v-bff6f798="">籍贯:四川阿坝</div><div class="m-btn pink" data-v-bff6f798="">体型:瘦长</div><div class="m-btn pink" data-v-bff6f798="">不吸烟</div><div class="m-btn pink" data-v-bff6f798="">稍微喝一点酒</div><div class="m-btn pink" data-v-bff6f798="">和家人同住</div><div class="m-btn pink" data-v-bff6f798="">未买车</div><div class="m-btn pink" data-v-bff6f798="">有孩子且住在一起</div><div class="m-btn pink" data-v-bff6f798="">是否想要孩子:视情况而定</div><div class="m-btn pink" data-v-bff6f798="">何时结婚:时机成熟就结婚</div></div></div>`
	submatch := compile.FindAllStringSubmatch(str, -1)
	fmt.Println(submatch)
	for i, item := range submatch {
		fmt.Println(i, item)
	}*/

	/*var compileNum = regexp.MustCompile("[0-9]+")
	fmt.Println(compileNum.FindString("fdsjfdlj123dfsf"))*/
	url2 := "http://album.zhenai.com/u/1863338240"
	bytes, e := fetcher.Fetch(url2)
	//profile := model.Profile{}
	if e != nil {
		panic(e)
	}
	submatch := compile2.FindAllStringSubmatch(string(bytes) /*profileRe2*/, -1)
	//parseResult := ParseProfile(bytes, profile)
	fmt.Println(submatch)
	fmt.Println(string(bytes))
}
