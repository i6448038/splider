package helper

import "strings"

//将相对路径改为绝对路径
func ChangeToAbspath(urls []string, hostName string)[]string{
	var res []string
	for _, e := range urls{
		res = append(res, GetAbspath(e, hostName))
	}
	return res
}

func GetAbspath(url string, hostName string)string{
	if strings.HasPrefix(url, "https://"){
		return url
	}
	return hostName + url
}