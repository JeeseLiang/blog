package pkg

import "time"

type ServerSettings struct {
	RunMode      string
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettings struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	CharSet      string
	ParseTime    bool
	MaxOpenConns int
	MaxIdleConns int
}

// 读取指定key的配置
func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
