package question

import (
	."splider/models"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

//解析知乎最主要的问题页
func PaserZhihuQuestion(url string)(*Crawler, error){
	crawlerData := new(Crawler)
	body, err := goquery.NewDocument(url)

	if err != nil{
		return crawlerData, err
	}

	crawlerData.Url = url

	questionHeader := body.Find(".QuestionPage .QuestionHeader .QuestionHeader-content")
	headerSide := questionHeader.Find(".QuestionHeader-side")
	headerMain := questionHeader.Find(".QuestionHeader-main")

	crawlerData.AttentionCount, err = strconv.Atoi(headerSide.
	Find(".NumberBoard.QuestionFollowStatus-counts .NumberBoard-value").
		Text())
	if err != nil{
		return crawlerData, err
	}

	crawlerData.Title = headerMain.Find(".QuestionHeader-title").Text()

	//answerCount, _ := body.Find(".Question-main .Card").Attr("data-za-module-info")
	//crawlerData.AnswerCount = int(answerCount)

	crawlerData.From = ZHIHU

	var tags string

	headerMain.Find(".QuestionHeader-tags .QuestionHeader-topics .Tag.QuestionTopic .Popover div").
		Each(func(i int, selection *goquery.Selection) {

		if len(tags) == 0 {
			tags = selection.Text()
		}else{
			tags = tags + " " + selection.Text()
		}
	})
	crawlerData.Tags = tags

	crawlerData.Desc = headerMain.Find(".QuestionHeader-detail span").Text()
	return crawlerData, nil
}