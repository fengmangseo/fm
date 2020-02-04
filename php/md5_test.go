package php_test

import (
	"fmt"
	"github.com/fengmangseo/fm/php"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(php.Md516("123"))
}
