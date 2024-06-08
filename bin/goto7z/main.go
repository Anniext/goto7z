package goto7z

import (
	"anniext.natapp4.cc/xt/goto7z/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goto7z",
		Short: `A simple automatic decompression 7z program`,
		Run: func(_ *cobra.Command, _ []string) {

		},
	}

	instanceProfile *profile.Profile
)

func init() {
	// 初始化配置文件
	cobra.OnInitialize(initConfig)
}

func Execute() error {
	return rootCmd.Execute()
}

// initConfig 初始化配置文件
func initConfig() {
	viper.AutomaticEnv() // 自动匹配环境变量的值进行赋值
	var err error
	instanceProfile, err = profile.GetProfile()
	if err != nil {
		slog.Error("failed to get profile", err)
		return
	}
}

func main() {
	err := Execute()
	if err != nil {
		panic(err)
	}
}
