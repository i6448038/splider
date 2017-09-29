package main

import (
	"fmt"
	."splider/spider_lib"
	."splider/models"
	"net/http"
)



func main(){

	defer func(){
		if p := recover(); p != nil{
			fmt.Println(p)
		}
	}()

	channel := make(chan []*Crawler)

	go ZhiHuBianJi(channel)
	go ZhihuDayhot(channel)
	go ZhihuMonthlyhot(channel)
	go ZhihuTopic(channel)

	for i:=0; i < 28; i++{
		msg := <-channel
		for _, v := range msg{
			_, err := Engine.Insert(v)
			if err != nil{
				fmt.Println("插入数据有误", ":", err.Error())
				return
			}
		}
	}
}