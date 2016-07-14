package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"io/ioutil"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"splider/kernel"
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

//func dbConnection(){
//	db, err := gorm.Open("mysql", "root:@/urls?charset=utf8&parseTime=True&loc=Local")
//}

