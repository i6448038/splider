package kernal

import (
	"time"
	"math/rand"
	//"strings"
	//"fmt"
)

//判断数组中有无此元素
func InArray(elem string, array []string)(result bool) {
	result = false
	for _,value:=range array{
		if(value == elem){
			result = true
		}
	}
	return
}
//生成随机字符串
func GetRandomString(lens int)(result string){
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//resultArray:=make([]string, len(str))
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i < lens;i++{
		result = result + (string)(str[r.Intn(len(str))])
	}
	return
}

//把一个字符串顺序打乱
//func Random(str string)string{
//	var result [len(str)]string
//	index:=rand.Intn(len(str)-1)
//	rand.Perm([]int{1,2})
//	if result[index] == ""{
//
//	}
//	for key,value:=range str{
//
//	}
//}

//字符串数组 变为字符串
func ChangeToString(str []string)(result string){
	for i:=0;i < len(str)-1; i++{
		result = result+str[i]
	}
	return
}