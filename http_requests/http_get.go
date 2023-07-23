package http_requests

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"nmc_spider/log_manage"
)

//	如果你想持续不断地访问一个页面并复用连接，你应该使用一个长连接（keep-alive）而不是每次重新建立连接。在Go语言的http包中，默认情况下，使用http.Client进行请求时会自动启用长连接，也就是说，对同一个主机的多个请求会复用同一个TCP连接。这样可以避免每次请求都进行三次握手和TLS握手，提高请求的效率。
//
// 在使用http.Client进行请求时，每次获取到响应后，应该调用Body.Close()来释放响应体的资源。虽然这样做不会关闭TCP连接，但它会释放响应体的资源，使得它们可以被垃圾回收器回收，从而避免内存泄漏。因此，即使使用长连接，也应该在每次获取到响应后调用Body.Close()。
// 需要注意的是，如果你使用了HTTP/2协议，则每次请求的响应体都需要在读取完毕后关闭，以便释放资源。在HTTP/2中，每个请求和响应都是一个帧，而不是一个完整的TCP流，因此需要确保在读取完响应体后关闭响应体。在Go语言的http包中，这个过程是自动完成的，无需手动调用Body.Close()。
var logger = log_manage.FSLogger
var gCurCookieJar, _ = cookiejar.New(nil)
var HttpClient = &http.Client{
	CheckRedirect: nil,
	Jar:           gCurCookieJar,
}

func HttpGet(url, uuid string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Host", "www.nmc.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,zh-CN;q=0.8,zh;q=0.7,zh-TW;q=0.5,zh-HK;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Referer", "http://www.nmc.cn/")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	resp, err := HttpClient.Do(req)
	err = resp.Body.Close()
	if err != nil {
		logger.Errorf("%v Close %v", uuid, err)
	}
	if err != nil {
		logger.Errorf("%v HttpGet1 %v", uuid, err)
		var tempResp []byte
		return tempResp
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Errorf("%v HttpGet2 %v", uuid, err)
		}
		return body
	}

}
