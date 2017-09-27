package main

import (
	"fmt"
	."splider/spider_lib"
	."splider/models"
)



func main(){
	 channel := make(chan []*Crawler)

	go ZhiHuBianJi(channel)
	go ZhihuDayhot(channel)
	go ZhihuMonthlyhot(channel)
	//go send(channel,"1b")
	//go send(channel,"1c")

	for i:=0; i < 3; i++{
		msg := <-channel
		for _, v := range msg{
			num, err := Engine.Insert(v)
			if err != nil{
				fmt.Println("插入数据有误", ":", err.Error())
				return
			}
			fmt.Print("插入的数据为：", num)
		}
	}
}