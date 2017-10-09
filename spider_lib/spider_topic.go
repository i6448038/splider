package spider_lib

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	."splider/spider_lib/landing_page"
	"io/ioutil"
	"encoding/json"
	"strings"
	"fmt"
	"strconv"
	"time"
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

func ZhihuTopic(channel chan <- []*Crawler){

	for _, v := range topicMap{
		var data []*Crawler
		url := "https://www.zhihu.com/topic/"+ v +"/hot"
		urls := RemoveDuplicates(parser(url))

		for _ , url := range urls{
			crawlerData, err := PaserZhihuQuestion(url)
			if err == nil{
				data = append(data, crawlerData)
			}
		}
		channel <- data
	}

	for _, v := range topicSpecial{
		var data []*Crawler
		url := "https://www.zhihu.com/topic/"+ v +"/top-answers"
		urls := RemoveDuplicates(parser(url))
		for _ , url := range urls{
			crawlerData, err := PaserZhihuQuestion(url)
			if err == nil{
				data = append(data, crawlerData)
			}
		}
		channel <- data
	}


}

func parser(url string)[]string{
	fmt.Println(url)
	resp := Get(url)
	defer resp.Body.Close()

	body, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		panic(err)
	}

	var urlList []string
	feedItems := body.Find(".feed-item.feed-item-hook")

	itmes := feedItems.Find("h2 a")

	if len(itmes.Nodes) == 0 {
		time.Sleep(time.Minute)
		return parser(url)
	}

	itmes.
		Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")

		if isExist{
			urlList = append(urlList, url)
		}
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	})

	for i := 2; len(urlList) < 100; i++{

		if  inSpecialTopics(url){
			feedItems = nextSpecial19Page(strconv.Itoa(i), url)
		}else {
			feedItems = next6Page(url, feedItems)
		}

		feedItems.Find(".feed-item.feed-item-hook h2 a").Each(func(i int, selection *goquery.Selection) {
			url, isExist := selection.Attr("href")

			if isExist{
				urlList = append(urlList, url)
			}
		})
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	}

	//urlList = RemoveDuplicates(urlList)
	return urlList
}

//一般领域的翻页
func next6Page(url string, document *goquery.Selection)*goquery.Selection{

	offset, isExist := document.Last().Attr("data-score")

	if !isExist{
		panic("获取下一页出问题")
	}

	resp := Post(url, offset)

	defer resp.Body.Close()

	content, error := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if error != nil {
		panic(error)
	}


	type Items struct {
		R int `json:"r"`
		Msg []interface{} `json:"msg"`
	}

	e := new(Items)

	error = json.Unmarshal(content, e)

	if error != nil{
		fmt.Println(url, offset)
		fmt.Println(string(content))
		panic(error)
	}

	html, ok := e.Msg[1].(string)

	if !ok {
		panic("强制类型转换失败")
	}

	respBody, error := goquery.NewDocumentFromReader(strings.NewReader(html))

	if error != nil{
		panic(error)
	}

	return respBody.Find(".feed-item.feed-item-hook")
}

func nextSpecial19Page(page, url string) *goquery.Selection{
	resp := Get(url + "?page="+page)
	defer resp.Body.Close()
	body, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		panic(err)
	}
	return body.Find("#zh-topic-top-page-list")
}

func inSpecialTopics(url string)bool{
	ret := false
	for _, v := range topicSpecial{
		if(url == "https://www.zhihu.com/topic/"+ v +"/top-answers"){
			ret = true
		}
	}
	return ret
}
