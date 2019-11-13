package fm

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func CompressHtml(b string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	s, _ := m.String("text/html", b)
	return s
}
