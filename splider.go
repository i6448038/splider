package main

import (
	//"fmt"
	"github.com/PuerkitoBio/goquery"
	//"io/ioutil"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"splider/kernel"
	"errors"
	//"net/http"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"sync"
	"container/list"
	"time"
	"strconv"
)

var (
	URL list.List
	lock sync.RWMutex
	//URLChannel chan string
)

func init(){
	kernal.Parse()
	//URLChannel = make(chan string, 1024)
	URL.PushBack(kernal.Property["url"])
}

func main() {
	var poolCount int
	ThreadLoop:for e := URL.Front(); e != nil; e = e.Next() {
		num,_:=kernal.Property["goroutineNum"]
		goroutineNums ,err := strconv.Atoi(num)
		if err!=nil {
			goroutineNums = 20
		}
		if poolCount < goroutineNums {
			url,ok := (e.Value).(string)
			if ok {
				lock.Lock()
				URL.Remove(e)
				lock.Unlock()
				go getHtmlData(url)
			}
		}
		poolCount++
	}

	for URL.Front() != nil{
		time.Sleep(3*time.Second)
		goto ThreadLoop
	}
	time.Sleep(10*time.Second)
}


func getHtmlData(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		labelA := s.Find("a").Nodes
		for _,attributes:= range labelA{
			for _,attr := range attributes.Attr{
				if attr.Key == "href"{
					if(!strings.Contains(attr.Val, "javascript")&&attr.Val!="/"){
						lock.Lock()
						URL.PushBack((string)(attr.Val))
						lock.Unlock()
					}
				}
			}
		}
		switch kernal.GetResource() {
		case "img":downloadImg( kernal.GetPath(), s)
		}
	})

}

func downloadImg( path string, s *goquery.Selection){
	imgs := s.Find("img").Nodes
	for _,attributes:= range imgs{
		for _,attr := range attributes.Attr{
			if attr.Key == "src" && attr.Val != "true" && len(attr.Val) > 0{
				url := attr.Val
				if !strings.Contains(attr.Val, "http"){
					url = kernal.GetURL()+attr.Val
				}
				postfix:=strings.SplitAfter(attr.Val, "/")
				resp,error:=http.Get(url)
				if error != nil{
					panic(errors.New(error.Error()))
				}
				image,_:=ioutil.ReadAll(resp.Body)
				error = ioutil.WriteFile(path+"/"+postfix[len(postfix)-1], image, os.ModePerm)
				fmt.Println(path+"/"+postfix[len(postfix)-1])
				if error!=nil{
					panic(errors.New(error.Error()))
				}
			}
		}
	}
}

