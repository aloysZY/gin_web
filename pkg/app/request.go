// Package app 接口校验二次封装
package app

import (
	"strings"

	"github.com/aloysZy/gin_web/pkg/setting"
	"github.com/gin-gonic/gin"
	val "github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// ValidError 解析错误结构体
type ValidError struct {
	Key     string
	Message string
}

// 实现了 error 方法，如果一个类型实现了某个 interface 中的所有方法，那么编译器认为该类型实现了此 interface，认为他们是"一样"的
func (v *ValidError) Error() string {
	return v.Message
}

// ValidErrors 定义别名
type ValidErrors []*ValidError

// ErrorF 实现格式化输出方法
func (v ValidErrors) ErrorF() string {
	return strings.Join(v.Errors(), ",")
}

// Errors 循环错误，返回一个切片
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValid 封装后，进行解析入参，调用此函数,没有错误，返回 true 和 nil，否则返回 false,和错误
func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	var errs ValidErrors
	// 内部根据Content-Type去解析
	//	c.Bind(obj interface{})
	//	内部替你传递了一个binding.JSON，对象去解析
	//	c.BindJSON(obj interface{})
	//	解析哪一种绑定的类型，根据你的选择
	//	c.BindWith(obj interface{}, b binding.Binding)
	// 内部根据Content-Type去解析
	// c.ShouldBind(obj interface{})
	// 内部替你传递了一个binding.JSON，对象去解析
	// c.ShouldBindJSON(obj interface{})
	// 解析哪一种绑定的类型，根据你的选择
	// c.ShouldBindWith(obj interface{}, b binding.Binding)
	// 注意：Shouldxxx和bindxxx区别就是bindxxx会在head中添加400的返回信息，而Shouldxxx不会
	err := c.ShouldBind(v)
	if err != nil {
		zap.L().Error("c.ShouldBind error: ", zap.Error(err))
		// 判断是不是ValidationErrors错误类型
		vErrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}
		// 		翻译器修改为全局的了，不安会并发不安全，测试一下
		// 取出翻译器
		// v := c.Value("trans")
		// 断言翻译器类型
		// trans, _ := v.(ut.Translator)
		// 否则进行翻译
		// for key, value := range vErrs.Translate(trans) {

		for key, value := range vErrs.Translate(setting.Trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
