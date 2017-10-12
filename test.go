package main

import (
	"splider/config"
)

func main(){
	config.Loggers["zhihu_access"].Println("出现错误！")
}
