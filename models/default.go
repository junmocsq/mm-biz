package models

import (
	"log"

	"github.com/junmo/mm-biz/models/goods"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&goods.Goods{}, &goods.Category{}, &goods.Brand{}, &goods.Company{})
	db.Migrator().AlterColumn(&goods.Company{}, "Name")
	db.Migrator().AlterColumn(&goods.Company{}, "ParentID")
}
