package setting

// viper 初始化
import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	// vp.SetConfigFile("./configs/config.yaml")
	// vp.WatchConfig()
	// vp.OnConfigChange(func(in fsnotify.Event) {
	// 	log.Println("夭寿啦~配置文件被人修改啦...")
	// 	vp.Unmarshal(&conf.Config)
	// })
	err := vp.ReadInConfig()
	if err != nil {
		// panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
		return nil, err
	}
	// if err := viper.Unmarshal(&conf.Config); err != nil {
	// 	panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	// }
	// log.Println("init viper success")
	return &Setting{vp}, nil
}
