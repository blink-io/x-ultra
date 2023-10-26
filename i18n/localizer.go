package i18n

func PluralCount(pluralCount interface{}) LOption {
	return func(config *LocalizeConfig) {
		config.PluralCount = pluralCount
	}
}

// L defines Localizer wrapper function for translation
func L(loc *Localizer) T {
	return func(messageID string, ops ...LOption) string {
		if loc != nil {
			conf := &LocalizeConfig{
				MessageID: messageID,
				Funcs:     fm,
			}
			for _, o := range ops {
				o(conf)
			}
			if s, err := loc.Localize(conf); err == nil {
				return s
			}
		}
		return messageID
	}
}

func (d MD) O() LOption {
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}

func D(d map[string]any) LOption {
	return func(c *LocalizeConfig) {
		c.TemplateData = d
	}
}
