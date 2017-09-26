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
	Desc string `xorm:"TEXT"`
	Img string `xorm:"null"`
	Tags string `xorm:"null"`
	AnswerCount int `xorm:"null"`
	AttentionCount int `xorm:"null"`
	PageView int `xorm:"null"`
	Origin int `xorm:"TINYINT"`
	CreatedAt time.Time `xorm:"TIMESTAMP created"`
}


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
