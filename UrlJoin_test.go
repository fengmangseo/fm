package fm_test

import (
	"github.com/fengmangseo/fm"
	"testing"
)

func TestUrlJoin(t *testing.T) {
	tophost, e := fm.UrlJoin("https://12423423.5346.a.com.cn/124/12/423/5/346/45//68", "../../1233.png")
	t.Error(tophost, e)
}
