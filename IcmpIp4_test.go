package fm

import (
	"fmt"
	"testing"
)

func TestIcmpIp4(t *testing.T) {
	e, ip := HostAllParsing("http://www.ruan8.cn/")
	fmt.Println(e, ip)
}
