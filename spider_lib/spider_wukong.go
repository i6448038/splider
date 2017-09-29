package spider_lib

import (
	."splider/models"
	."splider/helper"
	."splider/spider_lib/landing_page"
	"time"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type wukongResp struct {
	ErrorNo int
	LoginStatus int
	ErrTips string
	TotalNumber int
	HasMore bool
	HasMoreToRefresh bool
	ConcernData interface{}
	Data []respData `json:"data"`
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




func WukongList(channel chan <- []*Crawler){

	nowTimestamp := time.Now().UnixNano()/1e6

	url := "https://www.wukong.com/wenda/web/nativefeed/brow/?concern_id=6300775428692904450&t=" + strconv.FormatInt(nowTimestamp, 10)

	resp := Get(url)

	respJson, err:= ioutil.ReadAll(resp.Body)

	if err != nil{
		panic(err)
	}

	respContent := new(wukongResp)

	err = json.Unmarshal(respJson, respContent)

	if err != nil{
		panic(err.Error())
	}

	var urlList []string

	for _, v := range respContent.Data{
		if len(v.Question.Qid) > 0{
			urlList = append(urlList, "https://www.wukong.com/question/"  + v.Question.Qid + "/")
		}
	}


	var data []*Crawler

	for _, url := range ChangeToAbspath(urlList, "https://www.wukong.com"){
		crawlerData, err := PaserWukongQuestion(url)
		if err == nil{
			data = append(data, crawlerData)
		}
	}

	channel <- data
}
