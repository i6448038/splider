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
	//"sync"
)

var (
	URL []string
	URLChannel chan string
)

func init(){
	kernal.Parse()
	URLChannel = make(chan string, 1024)
}

func main() {
	for {
		//通过主URL获得众多子URL
		URL = append(URL, kernal.GetURL())
		for key, value := range URL {
			URL = append(URL[:key])
			go getHtmlData(value)
		}
		for _, value := range URLChannel {
			resp, ok := <-value
			if ok{
				URL = append(URL, resp)
			}

		}
		if len(URL) == 0 {
			break
		}
	}
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
						URLChannel <- (string)(attr.Val)

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

