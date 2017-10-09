package spider_lib

import (
	"github.com/PuerkitoBio/goquery"
	."splider/models"
	."splider/helper"
	."splider/spider_lib/landing_page"
	"strings"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)


func ZhiHuBianJi(channel chan <- []*Crawler){

	doc, err := goquery.NewDocument("https://www.zhihu.com/explore/recommendations")

	if err != nil{
		panic(err.Error())
	}

	var urlList []string
	doc.Find("#zh-recommend-list-full .zh-general-list .zm-item h2 a").
		Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		if isExist{
			urlList = append(urlList, url)
		}
		urlList = FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com"))
	})

	var data []*Crawler

	for i := 1; len(urlList) < 60; i++{
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
		urlList = FilterZhihuURLs(ChangeToAbspath(urlList, "https://www.zhihu.com"))
	}


	for _, url := range urlList{
		crawlerData, err := PaserZhihuQuestion(url)
		if err == nil{
			data = append(data, crawlerData)
		}
	}


	channel <- data
}

func nextBianjiPage(offset string, limit string)*goquery.Selection{
	ht := &http.Client{}
	resp, err := ht.Post("https://www.zhihu.com/node/ExploreRecommendListV2", "application/x-www-form-urlencoded",
		strings.NewReader(url.Values{"method":{"next"}, "params":{`{"limit":`+limit + `,"offset":` + offset + `}`}}.Encode()))
	if err != nil {
		panic(err)
	}

	content, error := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if error != nil {
		panic(error)
	}


	type Items struct {
		R int `json:"r"`
		Msg []interface{} `json:"msg"`
	}

	e := new(Items)

	error = json.Unmarshal(content, e)

	if error != nil{
		fmt.Println(string(content))
		panic(error)
	}

	html := ""

	for _, v := range e.Msg{
		msg, ok := v.(string)
		if ok{
			html = html + "\n" + msg
		}else{
			panic("强制类型转换失败")
		}
	}

	respBody, error := goquery.NewDocumentFromReader(strings.NewReader(html))

	if error != nil{
		panic(error)
	}

	return respBody.Find(".zm-item h2 a")
}


