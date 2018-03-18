package css

import (
	"log"
	"sync"
)

var (
	rules = &sync.Map{}

	STYLE_RULE     RuleType
	CHARSET_RULE   RuleType
	IMPORT_RULE    RuleType
	MEDIA_RULE     RuleType
	FONT_FACE_RULE RuleType
	PAGE_RULE      RuleType
)

type RuleType struct {
	name string
}

func (rt RuleType) Text() string {
	return rt.name
}

func init() {
	ruleTypeNames := []string{
		"@font-feature-values",
		"@keyframes",
		"@viewport",
		"@namespace",
		"@supports",
	}

	for _, n := range ruleTypeNames {
		AddNewType(n)
	}

	STYLE_RULE = RuleType{}
	CHARSET_RULE = AddNewType("@charset")
	IMPORT_RULE = AddNewType("@import")
	MEDIA_RULE = AddNewType("@media")
	FONT_FACE_RULE = AddNewType("@font-face")
	PAGE_RULE = AddNewType("@page")
}

type CSSRule struct {
	Type  RuleType
	Style CSSStyleRule
	Rules []*CSSRule
}

func NewRule(ruleType RuleType) *CSSRule {
	r := &CSSRule{
		Type: ruleType,
	}
	r.Style.Styles = make(map[string]*CSSStyleDeclaration)
	r.Rules = make([]*CSSRule, 0)
	return r
}

func AddNewType(name string) RuleType {
	r := RuleType{name: name}
	rules.Store(name, r)
	return r
}

func Get(name string) RuleType {
	if i, ok := rules.Load(name); ok {
		return i.(RuleType)
	}

	log.Printf("CSS request not existing rule: %s", name)
	return AddNewType(name)
}
