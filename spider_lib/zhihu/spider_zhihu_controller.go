package zhihu

import (
	"splider/config"
	."splider/models"
)

func ParseZhihuTopic(channel chan <- []*Crawler){
	datas, err := ZhiHuBianJi()

	if err != nil{
		config.Loggers["zhihu_error"].Println(err.Error())
	}
	channel <-datas


	datas, err = ZhihuDayhot()
	if err != nil{
		config.Loggers["zhihu_error"].Println(err.Error())
	}
	channel <-datas

	datas, err = ZhihuMonthlyhot()
	if err != nil{
		config.Loggers["zhihu_error"].Println(err.Error())
	}
	channel <-datas


	for _, v := range TopicMap{
		url := "https://www.zhihu.com/topic/"+ v +"/hot"
		datas, err := ZhihuTopic(url)
		if err != nil{
			config.Loggers["zhihu_error"].Println(err.Error())
		}
		channel <-datas
	}

	for _, v := range TopicSpecial{
		url := "https://www.zhihu.com/topic/"+ v +"/top-answers"
		datas, err := ZhihuTopic(url)
		if err != nil{
			config.Loggers["zhihu_error"].Println(err.Error())
		}
		channel <-datas
	}
}
