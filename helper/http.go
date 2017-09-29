package helper

import (
	"net/http"
	"strings"
	. "net/url"
)

func login(req *http.Request)*http.Request{
	req.AddCookie(_utma)
	req.AddCookie(_utmb)
	req.AddCookie(_utmc)
	req.AddCookie(_utmv)
	req.AddCookie(_utmz)
	req.AddCookie(_zap)
	req.AddCookie(cap_id)
	req.AddCookie(d_c0)
	req.AddCookie(q_c1)
	req.AddCookie(r_cap_id)
	req.AddCookie(z_c0)
	req.AddCookie(_xsrf)
	req.AddCookie(l_n_c)

	return req
}




func Post(url, offset string) *http.Response{
	ht := &http.Client{}
	resp, err := ht.Post(url, "application/x-www-form-urlencoded", strings.NewReader(Values{"start":{"0"}, "offset":{offset}}.Encode()))
	if err != nil {
		panic(err)
	}
	//req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36`)

	//resp, error := ht.Do(req)

	return resp
}


func Get(url string) *http.Response{
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36`)

	resp, error := http.DefaultClient.Do(req)


	if error != nil{
		panic(error)
	}


	return resp
}