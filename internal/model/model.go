package model

import (
	"blog/global"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=local",
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.Name,
		global.DatabaseSetting.CharSet,
		global.DatabaseSetting.ParseTime,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatalf("Setup DBEngine failed : %v\n", err)
		return nil, err
	}

	connPool, _ := db.DB()
	connPool.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)
	connPool.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)

	return db, nil
}
