package fm

import (
	"fmt"
	"github.com/fengmangseo/fm/php"
	"regexp"
	"testing"
)

func TestTakeChineseNumberFromString(t *testing.T) {
	fmt.Println(RangeChinaName())
	fmt.Println(regexp.MustCompile(`[css|js]$`).MatchString("/1242536.css"))
	//fmt.Println(TakeChineseNumberFromString("第万千五百四十七章"))
	//fmt.Println(DecodeToInt64("第千五百四十七章"))
	php.ArrayCountValues([]string{"123", "23"})
}
