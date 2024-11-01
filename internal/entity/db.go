package entity

import (
	"github.com/avrahambenaram/crud-produtos-go/internal/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(configuration.MysqlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	DB = db
}
