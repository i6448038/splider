package helper

import (
	"net/http"
	"time"
)

var timezone, _ = time.LoadLocation("Asia/Chongqing")

var _utma  = &http.Cookie{
	Name:"_utma",
	Value:"51854390.1158731505.1506597284.1506597284.1506597284.1",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2019, 9, 28, 19, 14, 0, 0, timezone),
}

var _utmb  = &http.Cookie{
	Name:"_utmb",
	Value:"51854390.0.10.1506597284",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2019, 9, 28, 7, 44, 0, 0, timezone),
}

var _utmc  = &http.Cookie{
	Name:"_utmc",
	Value:"51854390",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2018, 9, 28, 7, 44, 0, 0, timezone),
}

var _utmv  = &http.Cookie{
	Name:"_utmv",
	Value:"51854390.100-1|2=registration_date=20131104=1^3=entry_date=20131104=1",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2019, 9, 28, 7, 14, 0, 0, timezone),
}

var _utmz  = &http.Cookie{
	Name:"_utmz",
	Value:"51854390.1506597284.1.1.utmcsr=zhihu.com|utmccn=(referral)|utmcmd=referral|utmcct=/topic/19550994/hot",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2018, 3, 30, 7, 14, 0, 0, timezone),
}

var _zap  = &http.Cookie{
	Name:"_zap",
	Value:"34bdc08a-60e1-4f79-ad38-5a00e8a3f021",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2017, 10, 28, 19, 14, 0, 0, timezone),
}

var cap_id  = &http.Cookie{
	Name:"cap_id",
	Value:`"NjczODgyOTgyZDRmNDEyZWEzMGNhY2ZiZWZlMDk5NGM=|1506597272|e02f6a4bfb9781b3144deafc059db9d26339a729"`,
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2017, 10, 28, 19, 14, 0, 0, timezone),
}

var d_c0  = &http.Cookie{
	Name:"d_c0",
	Value:`"ABACPjUgcgyPTqAQFNe-qu1cdaR9LEZLmVw=|1506597283"`,
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2020, 9, 27, 19, 14, 0, 0, timezone),
}

var q_c1  = &http.Cookie{
	Name:"q_c1",
	Value:"b6b4f681d6fa4785aceb04d5ec1ba870|1506597272000|1506597272000",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2020, 9, 27, 7, 14, 0, 0, timezone),
}

var r_cap_id  = &http.Cookie{
	Name:"r_cap_id",
	Value:`"YWUzNzMyN2RhMzlhNGRhNThlNWM3NDdmOWQ3MzM5Y2Q=|1506597272|632f198b0b105088ec2166eddff034ef8dc6831b"`,
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2017, 10, 28, 19, 14, 0, 0, timezone),
}

var z_c0  = &http.Cookie{
	Name:"z_c0",
	Value:`"MS4xMUo4ZkFBQUFBQUFYQUFBQVlRSlZUYUZtOUZtbGRmUFFrNDNRVVFPY1E0cWRFM3Y1Qy1sQjl3PT0=|1506597281|c5d31af2fe7ce353e7b503af3525a2556a5624a2"`,
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2017, 10, 28, 19, 14, 0, 0, timezone),
}

var _xsrf  = &http.Cookie{
	Name:"_xsrf",
	Value:`9ddeb5694fd3d1af5de486f5c30adc40`,
	Path:"/",
	Domain:"www.zhihu.com",
	Expires:time.Date(2017, 10, 28, 19, 14, 0, 0, timezone),
}

var l_n_c  = &http.Cookie{
	Name:"l_n_c",
	Value:"1",
	Path:"/",
	Domain:".zhihu.com",
	Expires:time.Date(2018, 9, 28, 23, 9, 0, 0, timezone),
}


