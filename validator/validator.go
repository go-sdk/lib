package validator

import (
	enL "github.com/go-playground/locales/en"
	zhL "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enT "github.com/go-playground/validator/v10/translations/en"
	zhT "github.com/go-playground/validator/v10/translations/zh"
)

var (
	V *validator.Validate

	ZH ut.Translator
	EN ut.Translator
)

func init() {
	V = validator.New()
	zh := zhL.New()
	en := enL.New()
	uti := ut.New(zh, en)
	ZH, _ = uti.GetTranslator(zh.Locale())
	EN, _ = uti.GetTranslator(en.Locale())
	_ = zhT.RegisterDefaultTranslations(V, ZH)
	_ = enT.RegisterDefaultTranslations(V, EN)
}
