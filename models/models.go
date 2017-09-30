package models

import (
	"github.com/go-xorm/xorm"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"splider/config"
)

var Engine *xorm.Engine

type Crawler struct {
	Id  int  `xorm:"autoincr"`
	Url string `xorm:"null"`
	Title string `xorm:"null"`
	Desc string `xorm:"VARCHAR(3000) null"`
	Img []string `xorm:"null"`
	Tags string `xorm:"null"`
	AnswerCount int `xorm:"null"`
	AttentionCount int `xorm:"null"`
	Ext string `xorm:"VARCHAR(3000) null"`
	Pv int `xorm:"null"`
	From int `xorm:"TINYINT(1)"`
	Ctime time.Time `xorm:"TIMESTAMP created"`
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
	err = engine.Sync2(new(Crawler))
	if err != nil{
		panic(err)
	}
	Engine = engine
}
