package zhihu

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"splider/config"
	"time"
)


func ZhiHuBianJi(channel chan <- []*Crawler){
	client := &http.Client{}
	resp, _ := client.Get("https://www.zhihu.com/explore/recommendations")
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil{
		config.Loggers["zhihu_error"].Println("编辑推荐 刚启动协程就出现错误，协程关闭: ", err.Error())
		return
	}

	var urlList []string
	doc.Find("#zh-recommend-list-full .zh-general-list .zm-item h2 a").
		Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		if isExist{
			urlList = append(urlList, url)
		}
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	})

	var data []*Crawler

	for i := 1; len(urlList) < 100; i++{
		offset := ""
		if(i >= 4){
			offset = strconv.Itoa(i * 20 - 1)
		}else{
			offset = strconv.Itoa(i * 20)
		}
		nextBianjiPage(offset, "20").Each(func(i int, selection *goquery.Selection) {
			url, isExist := selection.Attr("href")
			if isExist{
				urlList = append(urlList, url)
			}
		})
		urlList = RemoveDuplicates(FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com")))
	}


	for _, url := range urlList{
		if err == nil{
			data = append(data, PaserZhihuQuestion(url))
		}
	}


	channel <- data
}

func nextBianjiPage(offset string, limit string)*goquery.Selection{
	ht := &http.Client{}
	resp, err := ht.Post("https://www.zhihu.com/node/ExploreRecommendListV2", "application/x-www-form-urlencoded",
		strings.NewReader(url.Values{"method":{"next"}, "params":{`{"limit":`+limit + `,"offset":` + offset + `}`}}.Encode()))
	defer resp.Body.Close()
	if err != nil {
		config.Loggers["zhihu_error"].Println("知乎编辑推荐 出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextBianjiPage(offset, limit)
	}

	content, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		config.Loggers["zhihu_error"].Println("知乎编辑推荐 出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextBianjiPage(offset, limit)
	}


	type Items struct {
		R int `json:"r"`
		Msg []interface{} `json:"msg"`
	}

	e := new(Items)

	error = json.Unmarshal(content, e)

	if error != nil{
		config.Loggers["zhihu_error"].Println("知乎编辑推荐 出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextBianjiPage(offset, limit)
	}

	html := ""

	for _, v := range e.Msg{
		msg, ok := v.(string)
		if ok{
			html = html + "\n" + msg
		}else{
			config.Loggers["zhihu_error"].Println("知乎编辑推荐 出现错误", err.Error(), "等待半分钟，重试！")
			time.Sleep(20 * time.Second)
			return nextBianjiPage(offset, limit)
		}
	}

	respBody, error := goquery.NewDocumentFromReader(strings.NewReader(html))

	if error != nil{
		config.Loggers["zhihu_error"].Println("知乎编辑推荐 出现错误", err.Error(), "等待半分钟，重试！")
		time.Sleep(20 * time.Second)
		return nextBianjiPage(offset, limit)
	}

	return respBody.Find(".zm-item h2 a")
}


