package kernal

import (
	"os"
	"bufio"
	"strings"
	"errors"
	//"fmt"
)

//conf 文件中必备的项
var (
	//PrimaryProperty = []string{"url"}
	Property   = map[string]string{
		"url":"",
	}
)

func GetURL() string{
	return Property["url"]
}

func Parse(){
	file,err:=os.Open("splider.conf")
	defer file.Close()
	if err !=nil  {
		panic(err)
	}
	buff:=bufio.NewReader(file)
	for{
		line,flag,_:=buff.ReadLine()
		content:=string(line)
		//如果conf文件的第一个字符不是#注释的话
		if(content[:1] != "#"){
			//防止同行出现#号的情况
			contentArray:=strings.Split(content, "=")
			//一行中如果用 = 分开后得到多个数据，则不符合规范。
			if(len(contentArray) == 2){
				_, ok:=Property[contentArray[0]]
				if ok{
					Property[contentArray[0]] = contentArray[1]
				}
			}else {
				panic(errors.New("conf 文件格式不正确"))
			}
			if flag == false{
				break
			}
		}

	}
}
