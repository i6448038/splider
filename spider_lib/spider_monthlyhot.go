package spider_lib

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/spider_lib/landing_page"
	."splider/helper"
)


func ZhihuMonthlyhot(channel chan <- []*Crawler){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#monthly-hot")

	if err != nil{
		panic(err.Error())
	}

	var urlList []string
	doc.Find("[data-type='monthly'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")

		if isExist{
			urlList = append(urlList, url)
		}
	})

	urlList = nextMonthPage("15", nextMonthPage("10", nextMonthPage("5", urlList)))

	var data []*Crawler

	for _, url := range FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")){
		crawlerData, err := PaserZhihuQuestion(url)
		if err == nil{
			data = append(data, crawlerData)
		}
	}

	channel <- data
}

func nextMonthPage(offset string, data []string)[]string{
	doc, err := goquery.NewDocument(`https://www.zhihu.com/node/ExploreAnswerListV2?params={"offset":` + offset + `,"type":"month"}`)

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