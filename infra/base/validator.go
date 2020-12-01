package base

import (
	"resk/infra"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Transtate() ut.Translator {
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate = validator.New()
	cn := zh.New()        // 中文翻译器创建
	uni := ut.New(cn, cn) // 通用翻译器创建UniversalTranslator
	var found bool
	translator, found = uni.GetTranslator("zh") // 获取通用中文翻译器
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator) // 向验证器注册翻译器
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Error("没有找到翻译器")
	}
}
