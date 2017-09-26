package main

import (
	"fmt"
	"splider/models"
)


func main(){
	data := new(models.Crawler)
	data.Url = "xxxx"
	data.AnswerCount = 1
	data.AttentionCount = 10
	data.Img = "zzzz"
	data.Desc = "dsafdsa"
	data.Tags = "xcvczxvcxzv"
	data.PageView = 123112321
	data.Origin = 1

	num, err := models.Engine.Insert(data)
	if err != nil{
		fmt.Println("插入数据有误", ":", err.Error())
		return
	}
	fmt.Print("插入的数据为：", num)
}
