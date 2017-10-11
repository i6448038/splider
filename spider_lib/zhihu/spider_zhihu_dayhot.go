package zhihu

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	"strconv"
	"time"
	"fmt"
)
func ZhihuDayhot(channel chan <- []*Crawler){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#daily-hot")

	if err != nil{
		panic(err.Error())
	}

	fmt.Println("开始抓每日热")
	var urlList []string
	doc.Find("[data-type='daily'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		fmt.Println(url)
		if isExist{
			urlList = append(urlList, url)
		}
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	})

	for i:=1; len(urlList) < 100; i++{
		offset := strconv.Itoa(i*5)
		urlList = RemoveDuplicates(append(urlList, FilterZhihuURLs(ChangeToAbspath(nextDayhotPage(offset,urlList), "https://www.zhihu.com"))...))
		fmt.Println("每日热list 长度", len(urlList))
	}


	var data []*Crawler

	for _, url := range urlList{
		data = append(data, PaserZhihuQuestion(url))
	}

	channel <- data

}

func nextDayhotPage(offset string, data []string)[]string{
	doc, err := goquery.NewDocument(`https://www.zhihu.com/node/ExploreAnswerListV2?params={"offset":` + offset + `,"type":"day"}`)

	if err != nil{
		fmt.Println("访问", "https://www.zhihu.com/node/ExploreAnswerListV2", "get", "正在等待一分钟")
		time.Sleep(time.Minute)
		return nextDayhotPage(offset, data)
	}

	doc.Find(".explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		data = append(data, url)

	})
	return data
}