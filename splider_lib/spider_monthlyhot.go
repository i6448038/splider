package splider_lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)


func ZhihuMonthlyhot(){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#monthly-hot")

	if err != nil{
		fmt.Println("连接错误!")
		return
	}

	doc.Find(".explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i)
		fmt.Println(selection.Text())
		fmt.Println(selection.Attr("href"))
	})
}