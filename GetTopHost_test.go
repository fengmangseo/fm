package fm_test

import (
	"github.com/fengmangseo/fm"
	"testing"
)

func TestGetTopHost(t *testing.T) {
	tophost, e := fm.GetTopHost("https://12423423.5346.a.com.cn/124")
	t.Error(tophost, e)

}
