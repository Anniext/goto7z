package profile

import (
	"github.com/spf13/viper"
)

// Profile 是启动主服务器的配置。
type Profile struct {
	Input   string `json:"input,omitempty"`                  // Input 是输入文件
	Output  string `json:"output,omitempty"`                 // Output 是输出文件
	Passwd  string `json:"passwd,omitempty"`                 // Passwd 是解压密码
	Mode    string `json:"mode,omitempty"`                   // Mode 是解压模式
	Version string `json:"version" json:"version,omitempty"` // Version 是服务器的当前版本
}

func GetProfile() (*Profile, error) {
	profile := Profile{}
	err := viper.Unmarshal(&profile)
	if err != nil {
		return nil, err
	}

	profile.Version = GetCurrentVersion(profile.Mode)

	return &profile, nil
}
