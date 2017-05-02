package css

import (
	"fmt"
	"sort"
	"strings"
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

	sort.Strings(decls)

	return fmt.Sprintf("%s { %s } ", sr.SelectorText, strings.Join(decls, ";"))
}
