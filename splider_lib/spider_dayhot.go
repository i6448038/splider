package splider_lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var data []string

func ZhihuDayhot(){

	defer func(){
		if p := recover(); p != nil{
			fmt.Println("存在错误！")
		}
	}()

	doc, err := goquery.NewDocument("https://www.zhihu.com/explore#daily-hot")

	if err != nil{
		panic(err.Error())
	}

	doc.Find("[data-type='daily'] .explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {

		url, _ := selection.Attr("href")
		data = append(data, selection.Text() + "  " + url)
	})

	data = nextPage("15", nextPage("10", nextPage("5", data)))

	fmt.Println(data)
}

func nextPage(offset string, data []string)[]string{
	doc, err := goquery.NewDocument(`https://wwws.zhihu.com/node/ExploreAnswerListV2?params={"offset":` + offset + `,"type":"day"}`)

	if err != nil{
		panic(err.Error())
		return []string{}
	}

	doc.Find(".explore-feed.feed-item h2 a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		data = append(data, selection.Text() + "  " + url)

	})
	return data
}