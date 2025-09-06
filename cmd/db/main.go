package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/junmo/mm-biz/models/goods"
	"github.com/junmo/mm-biz/models/other"
	"github.com/junmo/mm-biz/models/trade"
	"github.com/junmo/mm-biz/models/user"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	goodsModels := []interface{}{
		&goods.Brand{}, &goods.Category{}, &goods.Company{}, &goods.Goods{},
		&goods.Sale{}, &goods.Stock{}, &goods.Unit{},
	}
	db.Migrator().DropTable(goodsModels...)
	db.AutoMigrate(goodsModels...)

	otherModels := []interface{}{
		&other.Branch{}, &other.HeadQuarters{},
	}
	db.Migrator().DropTable(otherModels...)
	db.AutoMigrate(otherModels...)

	tradeModels := []interface{}{
		&trade.Retail{}, &trade.RetailGoods{},
		&trade.Wholesale{}, &trade.WholesaleGoods{},
	}
	db.Migrator().DropTable(tradeModels...)
	db.AutoMigrate(tradeModels...)

	userModels := []interface{}{
		&user.User{}, &user.Wholesaler{},
	}
	db.Migrator().DropTable(userModels...)
	db.AutoMigrate(userModels...)
	log.Println("Migration completed")
}
