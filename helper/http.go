package helper

import (
	"net/http"
	"strings"
	. "net/url"
)

func Post(url, offset string) *http.Response{
	ht := &http.Client{}
	resp, err := ht.Post(url, "application/x-www-form-urlencoded", strings.NewReader(Values{"start":{"0"}, "offset":{offset}}.Encode()))
	if err != nil {
		panic(err)
	}

	return resp
}


func Get(url string) *http.Response{
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, error := client.Do(req)


	if error != nil{
		panic(error)
	}

	return resp
}