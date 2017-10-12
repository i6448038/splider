package config

import (
	"github.com/Unknwon/goconfig"
	"path/filepath"
)

func loadConfigFile() *goconfig.ConfigFile{
	dirPath, _ := filepath.Abs("./")
	config , err := goconfig.LoadConfigFile(dirPath + "/config.ini")
	if err != nil{
		panic(err)
	}
	return config
}

