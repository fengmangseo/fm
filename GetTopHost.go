package fm

import (
	"net/url"
	"regexp"
	"strings"
)

//获取顶级域名
//From https://blog.csdn.net/u010071211/article/details/88690531
func GetTopHost(_url string) (tophost string, e error) {
	_url = strings.ToLower(_url)
	_url_parse, e := url.Parse(_url)
	if e != nil {
		return
	}
	_host_list := strings.Split(_url_parse.Hostname(), ".")
	_host_len := len(_host_list)
	matchString := regexp.MustCompile(`[\w].+\.(com|net|org|gov|edu)\.cn$`).MatchString(_url_parse.Host)
	if matchString && _host_len > 2 {
		tophost = _host_list[_host_len-3] + "." + _host_list[_host_len-2] + "." + _host_list[_host_len-1]
	} else {
		tophost = _host_list[_host_len-2] + "." + _host_list[_host_len-1]
	}
	return
}
