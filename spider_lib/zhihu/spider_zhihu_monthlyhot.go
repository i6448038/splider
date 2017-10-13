package zhihu

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	"strconv"
	"net/http"
	"splider/config"
	"time"
)


func ZhihuMonthlyhot()([]*FreeSpider, error){
	client := &http.Client{}
	resp, err := client.Get("https://www.zhihu.com/explore#monthly-hot")

	if err != nil{
		config.Loggers["zhihu_error"].Println("本月最热 刚启动协程就出现错误，协程关闭: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		config.Loggers["zhihu_error"].Println("本月最热 刚启动协程就出现错误，协程关闭: ", err.Error())
		return nil, err
	}

	var urlList []string
	doc.Find("[data-type='monthly'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")

		if isExist{
			urlList = append(urlList, url)
		}
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	})

	for i := 1; len(urlList) < 100; i++{
		time.Sleep(3 * time.Second)
		offset := strconv.Itoa(i*5)
		urlList = RemoveDuplicates(append(urlList, FilterZhihuURLs(ChangeToAbspath(nextMonthPage(offset,urlList), "https://www.zhihu.com"))...))
	}

	var data []*FreeSpider

	for _, url := range urlList{
		data = append(data, PaserZhihuQuestion(url))
	}

	return data, nil
}

func nextMonthPage(offset string, data []string)[]string{
	doc, err := goquery.NewDocument(`https://www.zhihu.com/node/ExploreAnswerListV2?params={"offset":` + offset + `,"type":"month"}`)

	if err != nil{
		config.Loggers["zhihu_error"].Println("解析知乎本月最热出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextMonthPage(offset, data)
	}

	doc.Find(".explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		data = append(data, url)

	})
	return data
}