package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"splider/helper"
)

func main()  {
	resp:= helper.Get("https://www.zhihu.com/question/66430136/answer/242147307")
	doc, _ := goquery.NewDocumentFromResponse(resp)
	fmt.Println(doc.Find(".QuestionPage .QuestionHeader .QuestionHeader-content .QuestionHeader-side .NumberBoard.QuestionFollowStatus-counts .NumberBoard-value").First().Text())
}
