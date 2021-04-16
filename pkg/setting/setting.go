package setting

import "github.com/spf13/viper"

/*
@Time    : 2021/4/16 8:46 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : setting.go
@Software: GoLand
*/
// read configs

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
