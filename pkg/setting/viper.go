// Package setting 初始化相关
package setting

// viper 初始化
import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var sections = make(map[string]any)

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
	s := &Setting{vp}
	// 热更新配置，这个项目启动初始化需要的参数是修改后没有效果的
	// 目前我尝试上传文件列表中的元素修改是可以的，端口和数据库修改没效果
	s.WatchSettingChange()
	return s, nil
}

// ReadSection 序列化到结构体
func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	// 配置文件读取后缓存
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

// ReloadAllSection 从新读取配置文件
func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		if err := s.ReadSection(k, v); err != nil {
			return err
		}
	}
	return nil
}

// WatchSettingChange 监听读取配置文件
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig() // 监听配置文件
		// 为Viper提供一个回调函数，以便在每次发生更改时运行
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("Config file changed:", in.Name)
			_ = s.ReloadAllSection()
		})
	}()
}
