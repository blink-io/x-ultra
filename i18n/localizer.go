package i18n

import (
	"text/template"
)

// Tr defines Localizer wrapper function for translation
func Tr(loc *Localizer) T {
	var fn T = func(messageID string, ops ...LOption) string {
		if loc != nil {
			cfg := &LocalizeConfig{
				MessageID: messageID,
				Funcs:     fm,
			}
			for _, o := range ops {
				o(cfg)
			}
			if s, err := loc.Localize(cfg); err == nil {
				return s
			}
		}
		return messageID
	}
	return fn
}

func PluralCount(pluralCount interface{}) LOption {
	return func(config *LocalizeConfig) {
		config.PluralCount = pluralCount
	}
}

func DefaultMessage(message *Message) LOption {
	return func(c *LocalizeConfig) {
		c.DefaultMessage = message
	}
}

func (d MD) O() LOption {
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}

func Funcs(funcMap template.FuncMap) LOption {
	return func(c *LocalizeConfig) {
		c.Funcs = funcMap
	}
}

func D(d map[string]any) LOption {
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}
