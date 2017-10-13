package models

import (
	"github.com/go-xorm/xorm"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"splider/config"
)

var Engine *xorm.Engine

type FreeSpider struct {
	Cid  int64  `xorm:"BIGINT(20) pk"`
	Url string `xorm:"null"`
	Title string `xorm:"null"`
	Desc string `xorm:"VARCHAR(6000) null"`
	Img []string `xorm:'[]'`
	Tags string `xorm:"null"`
	AnswerCount int `xorm:"null"`
	AttentionCount int `xorm:"null"`
	Ext string `xorm:"VARCHAR(3000) null"`
	Pv int `xorm:"null"`
	From int `xorm:"TINYINT(1)"`
	Ctime time.Time `xorm:"TIMESTAMP created"`
	Status int `xorm:"TINYINT(1)"`
}

const ZHIHU = 1
const WUKONG = 2


func init(){
	engine, err := xorm.NewEngine("mysql", config.DBconfig["user"] + ":" + config.DBconfig["pwd"] +
		"@tcp("+ config.DBconfig["local"] + ":" + config.DBconfig["port"] +")/" + config.DBconfig["db_name"])
	if err != nil{
		panic(err)
	}
	engine.SetMaxOpenConns(20)
	err = engine.Sync2(new(FreeSpider))
	if err != nil{
		panic(err)
	}
	Engine = engine
}

func SaveToMysql(datas []*FreeSpider){
	for _, data := range datas{
		crawler := new(FreeSpider)
		Engine.Where("url=?", data.Url).Get(crawler)
		if crawler.Url == ""{
			_, err := Engine.InsertOne(data)
			if err != nil{
				config.Loggers["zhihu_error"].Println("插入数据有误", ":", err.Error())
				config.Loggers["wukong_error"].Println("插入数据有误", ":", err.Error())

			}
		}
	}
}
