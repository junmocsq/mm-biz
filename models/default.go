package models

import (
	"github.com/junmocsq/jlib/dbcache"
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mm?charset=utf8mb4&parseTime=True&loc=Local"
	dbcache.RegisterDb(dsn, "mm", true)
	dbcache.RegisterCache(dbcache.NewLocalCache())
	//dbcache.RedisCacheInit("127.0.0.1", "6379", "")
	//dbcache.RegisterCache(NewRedisCache())
}
