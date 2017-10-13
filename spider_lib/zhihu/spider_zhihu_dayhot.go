package zhihu

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	"strconv"
	"time"
	"splider/config"
)
func ZhihuDayhot()([]*FreeSpider, error){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#daily-hot")

	if err != nil{
		config.Loggers["zhihu_error"].Println("今日最热 刚启动协程就出现错误，协程关闭: ", err.Error())
		return nil, err
	}

	var urlList []string
	doc.Find("[data-type='daily'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		if isExist{
			urlList = append(urlList, url)
		}
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	})

	for i:=1; len(urlList) < 100; i++{
		time.Sleep(3 * time.Second)
		offset := strconv.Itoa(i*5)
		urlList = RemoveDuplicates(append(urlList, FilterZhihuURLs(ChangeToAbspath(nextDayhotPage(offset,urlList), "https://www.zhihu.com"))...))
	}


	var data []*FreeSpider

	for _, url := range urlList{
		data = append(data, PaserZhihuQuestion(url))
	}

	return data, nil
}

func nextDayhotPage(offset string, data []string)[]string{
	doc, err := goquery.NewDocument(`https://www.zhihu.com/node/ExploreAnswerListV2?params={"offset":` + offset + `,"type":"day"}`)

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎今日最热出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(30 * time.Second)
		return nextDayhotPage(offset, data)
	}

	doc.Find(".explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		data = append(data, url)

	})
	return data
}