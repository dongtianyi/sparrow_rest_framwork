package settings

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// var myLogger = logging.GetLogger()

// 导出函数, 如果需要使用其他函数. 参考: https://github.com/spf13/viper/blob/master/viper.go
var (
	GetString   = viper.GetString   // 获取字符串
	GetBool     = viper.GetBool     // 获取 bool
	GetDuration = viper.GetDuration // 获取时间长度
	GetInt      = viper.GetInt      // 获取整形
	GetInt64    = viper.GetInt64    // 获取64位整形
)

func Init() {
	viper.SetConfigName("settings") // name of config file (without extension)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		// panic(fmt.Errorf("配置文件不存在: %s", err))
		fmt.Printf("配置文件不存在: %s", err)
		// myLogger.Error("当前文件夹没有找到 配置文件")
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

}

// func GetString(key string) string {
// 	val := viper.GetString(key)
// 	return val
// }
