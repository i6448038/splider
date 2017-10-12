package config


var DBconfig map[string]string

func init(){
	DBconfig = loadDBFile()
}


//加载配置文件的信息
func loadDBFile()map[string]string{

	config := loadConfigFile()
	ret := make(map[string]string)



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
