package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	viper *viper.Viper
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

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.viper.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
