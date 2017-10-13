package zhihu

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	"splider/config"
	"io/ioutil"
	"encoding/json"
	"strings"
	"strconv"
	"time"
	"net/http"
	."net/url"
)

var TopicMap = map[string]string{
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

var TopicSpecial = map[string]string{
	"投资":"19551404",
	"创业":"19550560",
}

func ZhihuTopic(url string)([]*FreeSpider, error){
	var data []*FreeSpider

	for _, url := range crawZhihuTopic(url){
		data = append(data, PaserZhihuQuestion(url))
	}
	return data, nil

}

func parser(url string)[]string{
	time.Sleep(3 * time.Second)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(30 * time.Second)
		return parser(url)
	}

	defer resp.Body.Close()

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(30 * time.Second)
		return parser(url)
	}

	body, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(30 * time.Second)
		return parser(url)
	}

	var urlList []string
	feedItems := body.Find(".feed-item.feed-item-hook")

	itmes := feedItems.Find("h2 a")

	if len(itmes.Nodes) == 0 {
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误, 可能让输入验证码", "等待半分钟，重试！")
		time.Sleep(30 * time.Second)
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
		time.Sleep(3 * time.Second)
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
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误, 可能让输入验证码", "等待半分钟，重试！")
		time.Sleep(30 * time.Second)
		return next6Page(url, document)
	}

	ht := &http.Client{}
	resp, err := ht.Post(url, "application/x-www-form-urlencoded", strings.NewReader(Values{"start":{"0"}, "offset":{offset}}.Encode()))
	if err != nil {
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return next6Page(url, document)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return next6Page(url, document)
	}


	type Items struct {
		R int `json:"r"`
		Msg []interface{} `json:"msg"`
	}

	e := new(Items)

	err = json.Unmarshal(content, e)

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return next6Page(url, document)
	}

	html, ok := e.Msg[1].(string)

	if !ok {
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现问题，可能是验证码，等待半分钟后重试！")
		time.Sleep(20 * time.Second)
		return next6Page(url, document)
	}

	respBody, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return next6Page(url, document)
	}

	return respBody.Find(".feed-item.feed-item-hook")
}

func nextSpecial19Page(page, url string) *goquery.Selection{
	client := &http.Client{}
	req, err := http.NewRequest("GET", url + "?page="+page, nil)
	if err != nil {
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextSpecial19Page(page, url)
	}

	resp, err := client.Do(req)

	if err != nil {
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextSpecial19Page(page, url)
	}

	defer resp.Body.Close()

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextSpecial19Page(page, url)
	}


	body, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎垂直专栏", url, "出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextSpecial19Page(page, url)
	}
	return body.Find("#zh-topic-top-page-list")
}

func inSpecialTopics(url string)bool{
	ret := false
	for _, v := range TopicSpecial{
		if(url == "https://www.zhihu.com/topic/"+ v +"/top-answers"){
			ret = true
		}
	}
	return ret
}


func crawZhihuTopic(url string)[]string{
	return RemoveDuplicates(parser(url))
}
