package landing_page

import (
	."splider/models"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"regexp"
	"strings"
	"fmt"
	."splider/helper"
	 "sync"
	"time"
)

var m *sync.Mutex
var ZhihuFlagCount = 0

func init(){
	m = new(sync.Mutex)
}


//过滤掉不符合要求的url
func FilterZhihuURLs(urls []string)[]string{
	var res []string
	for _, url := range urls{
		if regexp.MustCompile(`^https:\/\/www\.zhihu\.com\/question\/\d{1,}(\/answer\/\d{1,})?$`).MatchString(url){
			res = append(res, url)
		}
	}
	return res
}


//解析知乎最主要的问题页
func PaserZhihuQuestion(url string)(*Crawler, error){
	fmt.Println(url)
	crawlerData := new(Crawler)
	resp := Get(url)
	defer resp.Body.Close()
	body, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		return crawlerData, err
	}

	crawlerData.Url = url

	questionHeader := body.Find(".QuestionPage .QuestionHeader .QuestionHeader-content")
	headerSide := questionHeader.Find(".QuestionHeader-side")
	headerMain := questionHeader.Find(".QuestionHeader-main")

	crawlerData.AttentionCount, err = strconv.Atoi(headerSide.
	Find(".NumberBoard.QuestionFollowStatus-counts .NumberBoard-value").First().
		Text())
	if err != nil{
		return crawlerData, err
	}

	crawlerData.Title = headerMain.Find(".QuestionHeader-title").Text()

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
	crawlerData.AnswerCount, _ = strconv.Atoi(strings.TrimSuffix(strings.Replace(
		body.Find("#QuestionAnswers-answers").First().Find(".List-header h4 span").Text(), " ", "", -1),
			"个回答"))
	crawlerData.Pv, err = strconv.Atoi(headerSide.Find(".NumberBoard.QuestionFollowStatus-counts .NumberBoard-value").Last().Text())
	if err != nil{
		return crawlerData, err
	}
	var imgs = []string{}

	imgMess, isExist := body.Find("#data").Attr("data-state")

	if !isExist{
		panic("解析有问题")
	}

	reg := regexp.MustCompile(`\"editableDetail\":[\s]?\"([<\\>\s\t\w\d-=\\\"://\.])*,`)

	imgList, err := goquery.NewDocumentFromReader(
		strings.NewReader(
			strings.Replace(strings.TrimPrefix(strings.TrimSpace(reg.FindString(imgMess)), `"editableDetail":`), `\"`,`"`, -1)))

	if err != nil{
		panic(err)
	}

	imgList.Find("img").Each(func(i int, selection *goquery.Selection) {
		img, flag := selection.Attr("src")
		if flag {
			imgs = append(imgs, img)
		}
	})

	crawlerData.Img = imgs

	m.Lock()
	ZhihuFlagCount++
	if ZhihuFlagCount > 300{
		time.Sleep(10 * time.Second)
		ZhihuFlagCount = 0
	}
	m.Unlock()

	return crawlerData, nil
}