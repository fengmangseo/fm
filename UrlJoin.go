package fm

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func UrlJoin(_url string, s string) (s2 string, e error) {
	parse, e := url.Parse(_url)
	if e != nil {
		return
	}
	Path := parse.Path
	join := ""
	if strings.HasPrefix(s, "/") {
		join = s
	} else {
		join = regexp.MustCompile(`[\\+|\\]`).ReplaceAllString(filepath.Join(Path, s), "/")
	}

	s2 = parse.Scheme + "://" + parse.Host + join
	return
}
