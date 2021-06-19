package db

import (
	"blog/conf"
	"blog/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	var dbConfig = conf.Conf.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)

	Db, err := gorm.Open(conf.Conf.Db.Dialects, url)

	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(Db.Error)
	}

	Db.DB().SetMaxIdleConns(dbConfig.MaxIdle)
	Db.DB().SetMaxOpenConns(dbConfig.MaxOpen)
	logger := utils.Log()
	Db.SetLogger(logger)
	Db.LogMode(true)
	logger.Info("mysql connect success")

}
