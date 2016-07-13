package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"io"
	"bufio"
	//"io/ioutil"
	"strings"
	"os"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	value:=make(map[string]string)
	file,err:=os.Open("splider.conf")
	defer file.Close()
	if err !=nil  {
		panic(err)
	}
	buff:=bufio.NewReader(file)
	for{
		line,flag,_:=buff.ReadLine()
		valueArray:=strings.Split(string(line), "=")
		//一行中如果用 = 分开后得到多个数据，则不符合规范。
		if(len(valueArray) == 2){
			value[valueArray[0]]=valueArray[1]
		}else {
			panic(errors.New("conf 文件格式不正确"))
		}
		if flag == false{
			break
		}
	}
	if value["url"] == ""{
		panic(errors.New("url不合法！"))
	}
	getHtmlData(value["url"])
}

func getHtmlData(url string) {
	doc, err := goquery.NewDocument(strings.Trim(url, "\"\""))
	if err != nil {
		panic(err)
	}
	var urls []string
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		labelA := s.Find("a").Nodes
		for _,attributes:= range labelA{
			for _,attr := range attributes.Attr{
				if attr.Key == "href"{
					if(!strings.Contains(attr.Val, "javascript")&&attr.Val!="/"){
						urls = append(urls, attr.Val)
					}
				}
			}
		}
	})
	fmt.Println(urls)
}

func dbConnection(){
	db, err := gorm.Open("mysql", "root:@/urls?charset=utf8&parseTime=True&loc=Local")
}

