package config

import (
	"github.com/Unknwon/goconfig"
	"path/filepath"
)

var DBconfig map[string]string

func init(){
	DBconfig = loadFile()
}


//加载配置文件的信息
func loadFile()map[string]string{

	ret := make(map[string]string)
	dirPath, _ := filepath.Abs("./")
	config , err := goconfig.LoadConfigFile(dirPath + "/config/config.ini")
	if err != nil{
		panic(err)
	}

	user, err := config.GetValue("mysql", "user")
	if err != nil{
		panic(err)
	}
	ret["user"] = user

	pwd, err:= config.GetValue("mysql", "pwd")
	if err != nil{
		panic(err)
	}
	ret["pwd"] = pwd

	local, err:= config.GetValue("mysql", "local")
	if err != nil{
		panic(err)
	}
	ret["local"] = local

	port, err:= config.GetValue("mysql", "port")
	if err != nil{
		panic(err)
	}
	ret["port"] = port

	dbName, err:= config.GetValue("mysql", "db_name")
	if err != nil{
		panic(err)
	}
	ret["db_name"] = dbName

	return ret
}
