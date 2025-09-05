package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/junmo/mm-biz/models/goods"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Migrator().DropTable(&goods.Goods{}, &goods.Category{}, &goods.Brand{}, &goods.Company{})
	db.AutoMigrate(&goods.Goods{}, &goods.Category{}, &goods.Brand{}, &goods.Company{})

}
