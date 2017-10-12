package wukong

import (
	."splider/models"
	."splider/helper"
	"time"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"math/rand"
	"net/http"
	"splider/config"
)

type wukongResp struct {
	ErrorNo int
	LoginStatus int
	ErrTips string
	TotalNumber int
	HasMore bool
	HasMoreToRefresh bool
	ConcernData interface{}
}

type wukongNormalResp struct {
	wukongResp
	Data []respData `json:"data"`
}


type wukongRankResp struct {
	wukongResp
	RankData []respData `json:"rank_data"`
}

type wukongRankData struct {
	Answer interface{}
	Question question
}

type respData struct {
	Answer interface{}
	BehotTime int
	Cursor int
	Id int
	Question question `json:"question"`
	ShowAnswer bool
}

type question struct {
	ConcernTags interface{}
	CreateTime string
	DisplayStatus int
	FollowCount int
	GroupId int
	IsAnswer int
	IsFollow int
	IsSlave int
	Qid string `json:"qid"`
	Title string `json:"title"`
}

var domains = []string{
	"6300775428692904450",//热门
	"6215497896830175745",//娱乐
	"6215497726554016258",//体育
	"6215497898671475202",//汽车
	"6215497899594222081",//科技
	"6215497900164647426",//育儿
	"6215497899774577154",//美食
	"6215497897518041601",//数码
	"6215497898084272641",//时尚
	"6215847700051528193",//宠物
	"6215847700907166210",//收藏
	"6215497901804620289",//家居
	"6281512530493835777",//心理
	"6215497897710979586",//更多 文化
	"6215847700454181377",//更多 三农
	"6215497895248923137",//更多 健康
	"6215848044378720770",//更多 科学
	"6215497899027991042",//更多 游戏
	"6215497895852902913",//更多 动漫
	"6215497897312520705",//更多 教育
	"6215497899963320834",//更多 职场
	"6215497897899723265",//更多 旅游
	"6215497900554717698",//更多 电影
}

const (
	wukong_normal_url = "https://www.wukong.com/wenda/web/nativefeed/brow/?concern_id="
	wukong_rankhot_url = "https://www.wukong.com/wenda/web/hotrank/brow/?rank_type=0"
)

func WukongList(channel chan <- []*Crawler){

	var data []*Crawler
	wukongChannel := make(chan []string)

	//处理解析结构相同的领域
	for _,  domain := range domains{
		url := wukong_normal_url + domain + "&t=" +
			strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		go crawleWukongNormalTopic(domain, url, wukongChannel)
	}

	go crawleWukongSpecialTopic(wukongChannel)

	for i:=0; i <= len(domains); i++{
		urlList := <-wukongChannel
		for _, url := range ChangeToAbspath(urlList, "https://www.wukong.com"){
			crawlerData, err := PaserWukongQuestion(url)
			if err == nil{
				data = append(data, crawlerData)
			}
		}


		channel <- data
	}
}

//获取落地页地址
func getWukongLandingPageUrls(url string, rank bool)[]string{
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Close = true

	resp, err := client.Do(req)

	if err != nil{
		config.Loggers["wukong_error"].Println("访问", url, "出现问题," , err.Error(), "等待半分钟后重试!")
		time.Sleep(30 * time.Second)
		return getWukongLandingPageUrls(url, rank)
	}
	defer resp.Body.Close()


	respJson, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		config.Loggers["wukong_error"].Println("访问", url, "出现问题," , err.Error(), "等待半分钟后重试!")
		time.Sleep(30 * time.Second)
		return getWukongLandingPageUrls(url, rank)
	}


	var urlList []string

	if(!rank){
		respContent := new(wukongNormalResp)

		err = json.Unmarshal(respJson, respContent)

		if err != nil{
			time.Sleep(1 * time.Minute)
			return getWukongLandingPageUrls(url, rank)
		}

		for _, v := range respContent.Data{
			if len(v.Question.Qid) > 0{
				urlList = append(urlList, "https://www.wukong.com/question/"  + v.Question.Qid + "/")
			}
		}
	}else{
		respContent := new(wukongRankResp)

		err = json.Unmarshal(respJson, respContent)

		if err != nil{
			panic(err.Error())
		}

		for _, v := range respContent.RankData{
			if len(v.Question.Qid) > 0{
				urlList = append(urlList, "https://www.wukong.com/question/"  + v.Question.Qid + "/")
			}
		}
	}

	return urlList
}

//获取精选url，精选的规则和其他领域的不同
func getWokongRankUrl(page int) string{
	if page == 1 {
		return wukong_rankhot_url + "&t=" +strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	}else{
		return wukong_rankhot_url + "&t=" + strconv.FormatInt(time.Now().UnixNano()/1e6, 10) +
			"&new_offset=" + strconv.Itoa(page)
	}
}

//爬取悟空正常的数据
func crawleWukongNormalTopic(domain,url string, urls chan <-[]string){
	domainUrlList := RemoveDuplicates(getWukongLandingPageUrls(url, false))

	for len(domainUrlList) < 100 {
		url = wukong_normal_url + domain + "&t=" +
			strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + "&max_behot_time=" + strconv.FormatInt(time.Now().Add(time.Duration(rand.Int63n(120))  * time.Minute).Unix(), 10)

		domainUrlList = RemoveDuplicates(append(domainUrlList, getWukongLandingPageUrls(url, false)...))
	}
	urls <- domainUrlList
}

//爬取精选特殊的数据
func crawleWukongSpecialTopic(urls chan <-[]string){
	rankUrls := getWukongLandingPageUrls(getWokongRankUrl(1), true)
	rankUrls = append(rankUrls, getWukongLandingPageUrls(getWokongRankUrl(2), true)...)
	rankUrls = RemoveDuplicates(rankUrls)
	for i := 3; len(rankUrls) < 100; i++{
		rankUrls = RemoveDuplicates(append(rankUrls, getWukongLandingPageUrls(getWokongRankUrl(i), true)...))
	}
	urls <- rankUrls
}