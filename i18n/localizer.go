package i18n

import (
	"text/template"

	sprig "github.com/go-task/slim-sprig/v3"
)

// Tr defines Localizer wrapper function for translation
func Tr(loc *Localizer) T {
	var fn T = func(messageID string, ops ...LOption) string {
		if loc != nil {
			cfg := &LocalizeConfig{
				MessageID: messageID,
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

func SprigFuncs() LOption {
	return func(c *LocalizeConfig) {
		c.Funcs = sprig.FuncMap()
	}
}

func D(d map[string]any) LOption {
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}
