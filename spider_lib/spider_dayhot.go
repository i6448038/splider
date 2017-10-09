package spider_lib

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/spider_lib/landing_page"
	."splider/helper"
	"strconv"
)
func ZhihuDayhot(channel chan <- []*Crawler){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#daily-hot")

	if err != nil{
		panic(err.Error())
	}

	var urlList []string
	doc.Find("[data-type='daily'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		if isExist{
			urlList = append(urlList, url)
		}
		urlList = FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com"))
	})

	for i:=1; len(urlList) < 100; i++{
		offset := strconv.Itoa(i*5)
		urlList = append(urlList, FilterZhihuURLs(ChangeToAbspath(nextPage(offset,urlList), "https://www.zhihu.com"))...)
	}


	var data []*Crawler

	for _, url := range urlList{
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