package main

import (
	"fmt"
	."splider/spider_lib"
	."splider/models"
)

func main(){
	defer func(){
		if p := recover(); p != nil{
			fmt.Println(p)
		}
	}()

	channel := make(chan []*Crawler)

	go ZhiHuBianJi(channel)//1
	go ZhihuDayhot(channel)//1
	go ZhihuMonthlyhot(channel)//1
	go ZhihuTopic(channel)//28
	go WukongList(channel) //23

	for i := 0; i < 54; i++{
		_, err := Engine.Insert(<-channel)
		if err != nil{
			fmt.Println("插入数据有误", ":", err.Error())
			return
		}
	}
}