package config

import (
	"strings"
	"os"
	"time"
	"strconv"
	"log"
)

var logConfig map[string]string
var sectionList []string
var Loggers = make(map[string]*log.Logger)

func init(){
	logConfig = loadLogFile()
	Loggers = createLogers()
}


//加载配置文件的信息
func loadLogFile()map[string]string{
	config := loadConfigFile()

	ret := make(map[string]string)

	for _, v := range config.GetSectionList(){
		if topicName := strings.TrimSuffix(v,"_logs"); strings.HasSuffix(v, "_logs"){
			sectionList = append(sectionList, topicName)
			accessLog, err := config.GetValue(v, "access_log")
			if err != nil{
				panic(err)
			}
			ret[topicName + "_access"] = accessLog

			errorLog, err:= config.GetValue(v, "error_log")
			if err != nil{
				panic(err)
			}
			ret[topicName + "_error"] = errorLog
		}
	}

	return ret
}

func createLogers()map[string]*log.Logger{
	loggers := make(map[string]*log.Logger)
	now := time.Now()
	fileDate := strconv.Itoa(now.Year()) + strconv.Itoa(int(now.Month())) + strconv.Itoa(now.Day())
	for _, topicName := range sectionList{
		logName := strings.TrimSuffix(topicName,"_logs")
		f, error := os.Create(logConfig[logName + "_access"] + "_" + fileDate + ".log")
		if error != nil{
			panic(error)
		}
		loggers[strings.TrimSuffix(topicName, "_logs") + "_access"] = log.New(f,"", log.Ldate|log.Ltime)
		f, error = os.Create(logConfig[logName + "_error"] + "_" +fileDate + ".log")
		if error != nil{
			panic(error)
		}
		loggers[strings.TrimSuffix(topicName, "_logs") + "_error"] = log.New(f,"", log.Ldate|log.Ltime)
	}
	return  loggers
}

