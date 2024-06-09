package main

//go:generate go install anniext.natapp4.cc/xt/goto7z

import (
	"anniext.natapp4.cc/xt/goto7z/profile"
	"anniext.natapp4.cc/xt/goto7z/store"
	"fmt"
	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goto7z",
		Short: `A simple automatic decompression 7z program`,
		Run: func(_ *cobra.Command, _ []string) {

			store.InitBar(input)
			err := filepath.Walk(input, store.Visit)
			if err != nil {
				fmt.Printf("Error walking the path %q: %v\n", input, err)
				return
			}
			uiprogress.Stop()
		},
	}

	instanceProfile *profile.Profile
	input           string
	output          string
	passwd          string
	mode            string
)

func init() {
	// 初始化配置文件
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", `decompress input Path`)
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", `decompress output Path`)
	rootCmd.PersistentFlags().StringVarP(&passwd, "passwd", "p", "", `decompress password`)
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "7z", `decompression mode, which can be 7z or tar and zip`)

	err := viper.BindPFlag("input", rootCmd.PersistentFlags().Lookup("input"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("passwd", rootCmd.PersistentFlags().Lookup("passwd"))
	if err != nil {
		panic(err)
	}
	err = viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	if err != nil {
		panic(err)
	}

	workPath, _ := os.Getwd()
	viper.SetDefault("input", workPath)
	viper.SetDefault("output", workPath)
	viper.SetDefault("passwd", "costuan.com")
	viper.SetDefault("mode", "7z")
	viper.SetEnvPrefix("goto7z")
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
