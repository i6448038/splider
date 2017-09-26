package models

import (
	"github.com/go-xorm/xorm"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"splider/config"
)

var Engine *xorm.Engine

type Datas struct {
	Id  int32  `xorm:"autoincr"`
	Url string `xorm:"varchar(11) null"`
	CreatedAt time.Time `xorm:"TIMESTAMP null"`
}


func init(){
	engine, err := xorm.NewEngine("mysql", config.DBconfig["user"] + ":" + config.DBconfig["pwd"] +
		"@tcp("+ config.DBconfig["local"] + ":" + config.DBconfig["port"] +")/" + config.DBconfig["db_name"])
	if err != nil{
		panic(err)
	}
	engine.SetMaxOpenConns(20)
	err = engine.Sync2(new(Datas))
	if err != nil{
		panic(err)
	}
	Engine = engine
}
