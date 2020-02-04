package fm_test

import (
	"github.com/fengmangseo/fm"
	"testing"
)

func TestGetTopHost(t *testing.T) {

	tophost, e := fm.GetTopHost("5118.com")
	t.Error(tophost, e)
}
