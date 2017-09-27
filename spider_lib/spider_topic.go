package spider_lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var topicMap = map[string]string{
	"游戏":"19550994",
	"运动":"19552706",
	"互联网":"19550517",
	"艺术":"19550434",
	"阅读":"19550564",
	"美食":"19551137",
	"动漫":"19591985",
	"汽车":"19551915",
	"生活方式":"19555513",
	"教育":"19553176",
	"历史":"19551077",
	"文化":"19552266",
	"旅行":"19551556",
	"职业发展":"19554825",
	"足球":"19559052",
	"篮球":"19562832",
	"音乐":"19550453",
	"电影":"19550429",
	"法律":"19550874",
	"自然科学":"19553298",
	"设计":"19551557",
	"健康":"19550937",
	"商业":"19555457",
	"体育":"19554827",
	"科技":"19556664",
	"金融":"19609455",
}

var topicSpecial = map[string]string{
	"投资":"19551404",
	"创业":"19550560",
}

var ch = make(chan map[string][]string)

func ZhihuTopic(){

	var resultMap = make(map[string][]string)

	for k, v := range topicMap{
		url := "https://www.zhihu.com/topic/"+ v +"/hot"
		go parser(url, k)
	}

	for k, v := range topicSpecial{
		url := "https://www.zhihu.com/topic/"+ v +"/top-answers"
		go parser(url, k)
	}


	for i := 0; i < len(topicMap) + len(topicSpecial); i++{
		temp := make(map[string][]string)
		temp = <-ch
		for k , v := range temp{
			resultMap[k] = v
		}
	}

	fmt.Println(resultMap)
}

func parser(url, topicType string){
	doc, err := goquery.NewDocument(url)
	tempMap := make(map[string][]string)
	var dataArray []string
	if err != nil{
		//todo 打印log
		return
	}

	doc.Find(".feed-item.feed-item-hook h2 a").
		Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		dataArray = append(dataArray, selection.Text() + "\n" + url)
	})
	tempMap[topicType] = dataArray
	ch <- tempMap
}