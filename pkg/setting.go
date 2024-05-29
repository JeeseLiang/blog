package pkg

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 封装读取的配置文件
func NewSetting() (*Setting, error) {
	s := viper.New()

	// 添加配置文件的所在目录、文件名、后缀
	s.AddConfigPath("configs/")
	s.SetConfigName("config")
	s.SetConfigType("yaml")

	err := s.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{s}, nil
}
