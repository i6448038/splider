package main

import (
	"fmt"
	"splider/models"
)


func main(){
	data := new(models.Datas)
	data.Url = "xxxx"
	num, err := models.Engine.Insert(data)
	if err != nil{
		fmt.Println("插入数据有误", ":", err.Error())
		return
	}
	fmt.Print("插入的数据为：", num)
}
