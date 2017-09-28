package helper

import (
	"net/http"
	"strings"
	. "net/url"
)

func Post(url, offset string) *http.Response{
	req, err := http.NewRequest("POST", url, strings.NewReader(Values{"start":{"0"}, "offset":{offset}}.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36`)
	req.Header.Set("Referer", "url")
	resp, error := http.DefaultClient.Do(req)


	if error != nil{
		panic(error)
	}
	return resp
}
