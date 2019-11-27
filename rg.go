package fm

import "regexp"

const (
	RG_China_Zimu_Munber = "^[\u4e00-\u9fa5|0-9|a-z|A-Z+]$"
)

func StrFilter(src *string, rg_c string) {
	strn := ""
	for _, c := range *src {
		if regexp.MustCompile(rg_c).MatchString(string(c)) {
			strn += string(c)
		}
	}

	*src = strn
}
