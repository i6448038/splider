package spider_lib

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/spider_lib/question"
	."splider/helper"
	"fmt"
)
func ZhihuDayhot(channel chan <- []*Crawler){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#daily-hot")

	if err != nil{
		panic(err.Error())
	}

	var urlList []string
	doc.Find("[data-type='daily'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		fmt.Println(url)
		if isExist{
			urlList = append(urlList, url)
		}
	})

	urlList = nextPage("15", nextPage("10", nextPage("5", urlList)))

	var data []*Crawler

	for _, url := range FilterURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")){
		crawlerData, err := PaserZhihuQuestion(url)
		if err == nil{
			data = append(data, crawlerData)
		}
	}

	channel <- data

}

func nextPage(offset string, data []string)[]string{
	doc, err := goquery.NewDocument(`https://www.zhihu.com/node/ExploreAnswerListV2?params={"offset":` + offset + `,"type":"day"}`)

	if err != nil{
		panic(err.Error())
		return []string{}
	}

	doc.Find(".explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		data = append(data, url)

	})
	return data
}