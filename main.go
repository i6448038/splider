package main

import (
	"fmt"
	."splider/spider_lib/wukong"
	."splider/spider_lib/zhihu"
	."splider/models"
)

func main(){

	defer func(){
		if p := recover(); p != nil{
			fmt.Println(p)
		}
	}()

	channel := make(chan []*Crawler)

	go ParseZhihuTopic(channel)//5

	go WukongList(channel) //24

	for i := 0; i < 55; i++{
		SaveToMysql(<-channel)
	}

}