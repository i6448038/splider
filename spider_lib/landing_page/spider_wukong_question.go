package landing_page

import (
	. "splider/models"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"splider/helper"
	"fmt"
)

func PaserWukongQuestion(url string)(*Crawler, error){
	crawlerData := new(Crawler)
	body, err := goquery.NewDocument(url)

	if err != nil{
		return crawlerData, err
	}

	crawlerData.Url = url
	crawlerData.From = WUKONG

	question := body.Find(".question.question-single")

	questionMain := question.Find(".question-item")

	crawlerData.Title = questionMain.Find(".question-name").Text()

	fmt.Println(url)

	tags, isExist := questionMain.Find(`[itemprop="keywords"]`).Attr("content")

	if !isExist {
		crawlerData.Tags = tags
	}else {
		crawlerData.Tags = strings.Replace(tags,",","", -1)
	}

	var imgList []string

	questionMain.Find(".question-img-preview .image-box img").
		Each(func(i int, selection *goquery.Selection) {
		img, _:=selection.Attr("src")
		img = helper.GetAbspath(img, "https://www.wukong.com")
		imgList = append(imgList, img)
	})

	var img string

	for _, v := range imgList{
		if len(img) == 0{
			img = v
		}else{
			img = img + " " + v
		}

	}

	crawlerData.Img = img

	crawlerData.Desc = questionMain.Find(".question-text").Text()

	crawlerData.AttentionCount, err = strconv.Atoi(questionMain.Find("[data-node='followquestion'] .count").Text())

	if err != nil{
		panic(err)
	}

	crawlerData.AnswerCount, err = strconv.Atoi(strings.TrimSuffix(questionMain.Find(".answer-count-h span").Text(), "个回答"))

	if err != nil{
		panic(err)
	}

	return crawlerData, nil
}

