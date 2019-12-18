package fm

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func Pinyin(hans string) string {
	args := pinyin.NewArgs()
	args.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	return strings.Join(pinyin.LazyPinyin(hans, args), "")
}
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}
