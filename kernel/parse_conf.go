package kernal

import (
	"os"
	"strings"
	"io/ioutil"
	"errors"
)

//conf 文件中必备的项
var (
	Property   = map[string]string{
		"url":"",
		"resource":"img",
		"path":"Downloads",
	}
)
//获取索要访问的路径
func GetURL() string{
	return Property["url"]
}
//获取资源的类型
func GetResource() string{
	return Property["resource"]
}
//获取资源的存放路径
func GetPath() string{
	os.MkdirAll(GetRootPath()+Property["path"], 0777)
	return GetRootPath()+Property["path"]
}

func Parse(){
	file,err:=ioutil.ReadFile("splider.conf")
	if err !=nil  {
		panic(errors.New("读取文件有误！"))
	}
	fileContent:=string(file)
	lines:=strings.Split(fileContent, "\n")
	for i:=0;i<len(lines);i++ {
		content:=string(lines[i])
		//如果conf文件的第一个字符不是#注释的话
		if(len(content)>0 && content[:1] != "#" ){
			contentArray:=strings.Split(content, "=")
			//一行中如果用 = 分开后得到多个数据，则不符合规范。
			//防止同行出现#号的情况
			if(len(contentArray) == 2 && !strings.Contains(content, "#")){
				_, ok:=Property[contentArray[0]]
				if ok{
					value:=strings.Trim(contentArray[1], "\"\"\r")
					if contentArray[0] == "url"{
						if !strings.Contains(value, "http://"){
							value = "http://"+value
						}
					}
					Property[contentArray[0]] = value
				}
			}

		}
	}
}
