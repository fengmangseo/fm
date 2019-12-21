package fm

import (
	"github.com/axgle/pinyin"
	"strings"
)

func Pinyin(string2 string) string {
	return strings.ToLower(pinyin.Convert(string2))
}
