package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(`./config`)
	if err := viper.ReadInConfig(); err != nil {
		// 读取配置失败
		fmt.Printf("viper load config failed: %v\n", err)
		return err
	}
	// 支持热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("%s changed..\n", in.Name)
	})
	return
}
