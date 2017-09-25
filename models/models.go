package models

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var Engine *xorm.Engine

type Datas struct {
	Id  int32  `xorm:"autoincr"`
	Url string `xorm:"varchar(11) null"`
	CreatedAt time.Time `xorm:"TIMESTAMP null"`
}


func init(){
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:homestead@tcp(127.0.0.1:3306)/pholcus")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	Engine.SetMaxOpenConns(20)
	err = Engine.Sync2(new(Datas))
	if err != nil{
		fmt.Println(err.Error())
		return
	}
}
