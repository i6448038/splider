package spider_lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	"regexp"
	."splider/helper"
	."splider/spider_lib/question"
)


func ZhiHuBianJi(channel chan <- []*Crawler){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore/recommendations")

	if err != nil{
		fmt.Println("连接错误!")
		return
	}

	var urlList []string
	doc.Find("#zh-recommend-list-full .zh-general-list .zm-item h2 a").
		Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")

		if isExist{
			urlList = append(urlList, url)
		}
	})

	fmt.Println(ChangeToAbspath(urlList, "https://www.zhihu.com"))

	var data []*Crawler

	for _, url := range filterURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")){
		crawlerData, err := PaserZhihuQuestion(url)
		if err == nil{
			data = append(data, crawlerData)
		}
	}

	channel <- data
}
//过滤掉不符合要求的url
func filterURLs(urls []string)[]string{
	var res []string
	for _, url := range urls{
		if regexp.MustCompile(`^https:\/\/www\.zhihu\.com\/question\/\d{1,}\/answer\/\d{1,}$`).MatchString(url){
			res = append(res, url)
		}
	}
	return res
}


