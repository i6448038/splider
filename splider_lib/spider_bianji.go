package splider_lib

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)


func ZhiHuBianJi(){
	doc, err := goquery.NewDocument("https://www.zhihu.com/explore/recommendations")

	if err != nil{
		fmt.Println("连接错误!")
		return
	}

	doc.Find("#zh-recommend-list-full .zh-general-list .zm-item h2 a").
		Each(func(i int, selection *goquery.Selection) {
		fmt.Println(i)
		fmt.Println(selection.Text())
		fmt.Println(selection.Attr("href"))
	})

}