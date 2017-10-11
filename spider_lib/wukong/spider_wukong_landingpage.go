package wukong

import (
	. "splider/models"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"splider/helper"
	"fmt"
	"sync"
	"time"
)

var wukongMu *sync.Mutex
var wukongFlagCount = 0

func init(){
	wukongMu = new(sync.Mutex)
}

//解析悟空落地页
func PaserWukongQuestion(url string)(*Crawler, error){
	crawlerData := new(Crawler)
	body, err := goquery.NewDocument(url)

	if err != nil{
		return crawlerData, err
	}

	crawlerData.Url = strings.TrimSpace(url)
	crawlerData.From = WUKONG

	question := body.Find(".question.question-single")
	questionMain := question.Find(".question-item")
	crawlerData.Title = strings.TrimSpace(questionMain.Find(".question-name").Text())
	tags, isExist := questionMain.Find(`[itemprop="keywords"]`).Attr("content")

	if !isExist {
		crawlerData.Tags = tags
	}else {
		crawlerData.Tags = strings.Replace(tags,","," ", -1)
	}

	var imgList = []string{}

	questionMain.Find(".question-img-preview .image-box img").
		Each(func(i int, selection *goquery.Selection) {
		img, _:=selection.Attr("src")
		img = helper.GetAbspath(img, "https:")
		imgList = append(imgList, img)
	})

	crawlerData.Img = imgList
	crawlerData.Desc = strings.TrimSpace(questionMain.Find(".question-text").Text())
	crawlerData.AttentionCount, err = strconv.Atoi(questionMain.Find(".question-bottom [data-node='followquestion'] .count").Text())
	if err != nil{
		time.Sleep(10 * time.Second)
		return PaserWukongQuestion(url)
	}

	crawlerData.AnswerCount, err = strconv.Atoi(strings.TrimSuffix(questionMain.Find(".answer-count-h span").Text(), "个回答"))
	if err != nil{
		time.Sleep(10 * time.Second)
		return PaserWukongQuestion(url)
	}
	wukongMu.Lock()
	wukongFlagCount++
	if wukongFlagCount > 300{
		time.Sleep(10 * time.Second)
		wukongFlagCount = 0
	}
	wukongMu.Unlock()

	fmt.Println("成功爬取了url", url)

	return crawlerData, nil
}

