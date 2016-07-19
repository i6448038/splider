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
	doc, err := goquery.NewDocument(strings.Trim(url, "\"\""))
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
		resource,ok:= kernal.Property["resource"]
		if !ok {
			panic(errors.New("配置文件出错！"))
		}
		switch resource {
		case "img":
		}
	})

}

//func downloadImg(resourceType string, path string, s *goquery.Selection){
//	imgs := s.Find("img").Nodes
//	for _,attributes:= range imgs{
//		for _,attr := range attributes.Attr{
//			if attr.Key == "src"{
//				resp,_:=http.Get(attr.Val)
//				image,_:=ioutil.ReadAll(resp.Body)
//
//				ioutil.WriteFile(path+"/"+kernal.GetRandomString(10))
//			}
//		}
//	}
//}

