package base

import (
	"resk.com/infra"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	Check(validate)
	return validate
}

func Transtate() ut.Translator {
	Check(translator)
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

func ValidateStruct(s interface{}) (err error) {
	err = Validate().Struct(s)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			log.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				log.Error(e.Translate(Transtate()))
			}
		}
		return err
	}
	return nil
}
