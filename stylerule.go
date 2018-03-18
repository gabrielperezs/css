package css

import (
	"fmt"
	"strings"
)

var (
	newLine = string(byte('\n'))
	retLine = string(byte('\r'))
	tabChar = string(byte('\t'))
)

type CSSStyleRule struct {
	SelectorText string
	Styles       map[string]*CSSStyleDeclaration
}

func (sr *CSSStyleRule) Text() string {
	decls := make([]string, 0, len(sr.Styles))

	for _, s := range sr.Styles {
		decls = append(decls, s.Text())
	}

	sr.SelectorText = strings.TrimSpace(strings.Replace(sr.SelectorText, " ", " ", -1))
	sr.SelectorText = strings.Replace(sr.SelectorText, newLine, "", -1)
	sr.SelectorText = strings.Replace(sr.SelectorText, retLine, "", -1)
	sr.SelectorText = strings.Replace(sr.SelectorText, tabChar, "", -1)

	return fmt.Sprintf("%s{%s}", sr.SelectorText, strings.Join(decls, ";"))
}
