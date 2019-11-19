package fm

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//检查域名是否是泛解析
func HostAllParsing(host string) (e error, is_ok bool) {
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}
	tophost, e := GetTopHost(host)
	if e != nil {
		return
	}
	host = fmt.Sprintf("%d.%s", time.Now().UnixNano(), tophost)
	conn, e := net.Dial("ip4:icmp", host)
	if e != nil {
		return
	}
	defer conn.Close()
	is_ok = true
	return
}
