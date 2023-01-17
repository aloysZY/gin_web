package setting

// viper 初始化
import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting 初始化配置文件
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
