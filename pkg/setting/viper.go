// Package setting 初始化相关
package setting

// viper 初始化
import (
	"strings"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting 初始化配置文件
func NewSetting(configs ...string) (*Setting, error) {
	/*viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")               // 还可以在工作目录中查找配置*/
	vp := viper.New()
	for _, config := range configs {
		// 判断传入的路径是不是 yaml 结尾的，是就直接使用这个路径，不是就是默认是路径，要设置解析
		if strings.HasSuffix(config, ".yaml") {
			vp.SetConfigFile(config)
		} else {
			vp.AddConfigPath(config)
			vp.SetConfigName("config")
			vp.SetConfigType("yaml")
		}
	}
	// 读取配置
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

// ReadSection 序列化到结构体
func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
