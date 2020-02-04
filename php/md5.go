package php

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	return Md532(s)
}
func Md532(s string) string {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(s))
	Result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", Result)
}
func Md516(s string) string {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(s))
	Result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", Result)[8:24]
}
