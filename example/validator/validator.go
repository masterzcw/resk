package main

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

type User struct {
	FirstName string `validate:"required"` // 必填
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`  // >=0 且 <=130
	Email     string `validate:"required,email"` // 必填email
}

func main() {
	// 数据
	user := &User{
		FirstName: "firstName",
		LastName:  "lastName",
		Age:       240,
		Email:     "fl@163.com",
	}

	// 验证器 & 国际化
	validate := validator.New()
	cn := zh.New()                          // 中文翻译器创建
	uni := ut.New(cn, cn)                   // 通用翻译器创建UniversalTranslator
	trans, found := uni.GetTranslator("zh") // 获取通用中文翻译器
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, trans) // 向验证器注册翻译器
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("没有找到翻译器")
	}

	// 验证数据
	err := validate.Struct(user)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError) // err转换为(*validator.InvalidValidationError)类型, 表示无效的输入
		if ok {
			fmt.Println("InvalidValidationError:", err)
		}
		errs, ok := err.(validator.ValidationErrors) // 表示非法输入
		if ok {
			for _, err := range errs {
				fmt.Println("ValidationErrors:", err.Translate(trans)) // 把错误信息转换为中文
			}
		}
	}
}
