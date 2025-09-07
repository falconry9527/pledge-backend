package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// BindingValidator gin bind go-playground/validate
func BindingValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("IsPassword", IsPassword)                           //判断是否为合法密码
		_ = v.RegisterValidation("IsPhoneNumber", IsPhoneNumber)                     //检查手机号码字段是否合法
		_ = v.RegisterValidation("IsEmail", IsEmail)                                 //检查邮箱字段是否合法
		_ = v.RegisterValidation("CheckUserNicknameLength", CheckUserNicknameLength) //检查用户昵称长度是否合法
		_ = v.RegisterValidation("CheckUserAccount", CheckUserAccount)               //检查用户账号是否合法
		_ = v.RegisterValidation("OnlyOne", OnlyOne)                                 //字段唯一性约束
	}
	// 这些注册的方法，可以在 结构体中直接使用,例如
	// 1. 结构体标记
	// Password string `json:"password" binding:"required,IsPassword"` // 使用自定义验证
	// 2. 调用验证函数
	// validator.Verify(&req)

}
