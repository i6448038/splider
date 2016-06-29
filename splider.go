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
		if(len(valueArray) == 2){
			value[valueArray[0]]=valueArray[1]
		}else {
			panic(errors.New("conf 文件格式"))
		}
		if flag == false{
			break
		}
	}
	//getHtmlData("http://www.soufang.com")
}

func getHtmlData(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i)
		labelA := s.Find("a").Nodes
		var urls []string
		for _,value:= range labelA{
			for _,value1 := range value.Attr{
				if value1.Key == "href"{
					urls = append(urls, value1.Val)
				}
			}
		}
		fmt.Println(urls)
	})
}

