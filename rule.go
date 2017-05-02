package css

type RuleType int

const (
	STYLE_RULE RuleType = iota
	CHARSET_RULE
	IMPORT_RULE
	MEDIA_RULE
	FONT_FACE_RULE
	PAGE_RULE
	WEBKITKEYFRAMES
	KEYFRAMES
	MSVIEWPORT
	MOZDOCUMENT
)

var ruleTypeNames = map[RuleType]string{
	STYLE_RULE:      "",
	MEDIA_RULE:      "@media",
	CHARSET_RULE:    "@charset",
	IMPORT_RULE:     "@import",
	FONT_FACE_RULE:  "@font-face",
	PAGE_RULE:       "@page",
	WEBKITKEYFRAMES: "@-webkit-keyframes",
	KEYFRAMES:       "@keyframes",
	MSVIEWPORT:      "@-ms-viewport",
	MOZDOCUMENT:     "@-mod-document",
}

func (rt RuleType) Text() string {
	return ruleTypeNames[rt]
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
