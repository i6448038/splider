package zhihu

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	"strconv"
	"fmt"
	"net/http"
)


func ZhihuMonthlyhot(channel chan <- []*Crawler){
	client := &http.Client{}
	resp, _ := client.Get("https://www.zhihu.com/explore#monthly-hot")
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		panic(err)
	}

	if err != nil{
		panic(err.Error())
	}

	fmt.Println("开始抓每月热")
	var urlList []string
	doc.Find("[data-type='monthly'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")

		if isExist{
			urlList = append(urlList, url)
		}
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	})

	for i := 1; len(urlList) < 100; i++{
		offset := strconv.Itoa(i*5)
		urlList = RemoveDuplicates(append(urlList, FilterZhihuURLs(ChangeToAbspath(nextMonthPage(offset,urlList), "https://www.zhihu.com"))...))
		fmt.Println("每月热list 长度", len(urlList))
	}

	var data []*Crawler

	for _, url := range urlList{
		data = append(data, PaserZhihuQuestion(url))
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