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
	//"github.com/jinzhu/gorm"
	//"net/http"
	//"io/ioutil"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

var (
	URL []string
)

func init(){
	kernal.Parse()
}

func main() {
	getHtmlData(kernal.GetURL())
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
						URL = append(URL, attr.Val)
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
				postfix:=strings.SplitAfter(attr.Val, ".")
				resp,error:=http.Get(url)
				if error != nil{
					panic(errors.New(error.Error()))
				}
				image,_:=ioutil.ReadAll(resp.Body)
				error = ioutil.WriteFile(path+"/"+kernal.GetRandomString(10)+"."+postfix[len(postfix)-1], image, os.ModePerm)
				fmt.Println(path+"/"+kernal.GetRandomString(10)+"."+postfix[len(postfix)-1])
				if error!=nil{
					panic(errors.New(error.Error()))
				}
			}
		}
	}
}

