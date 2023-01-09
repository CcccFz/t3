package store

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"main/myTest/domain/entity"
)

var DB = MyMysql("t3")

var CACHE = MyCache(14)

func MyMysql(dbName string) (db *gorm.DB) {
	//_db := store.NewMysql("edu")(&store.MysqlConfig{Host: "192.168.50.16", Password: "z6skqQJrf",IsLog: false})
	_db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf(
				"root:%s@%s(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				"z6skqQJrf", "tcp", "192.168.50.16", dbName,
			),
			DefaultStringSize:        256,
			DisableDatetimePrecision: true,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		log.Fatal().Err(err).Str("host", "192.168.50.16").Str("db", dbName).Msg("mysql数据库连接失败")
	}
	db = _db.Debug()

	return
	//test := new(Test)
	//if err:= _db.AutoMigrate(test); err != nil {fmt.Println("error")}

	//newTest.Name = "Duome"
	//db.Create(&Test{Name: "Mama"})
	//test.Id = 1
	//test.Name = "Papa"
	//db.Model(test).Updates(test)
	////db.Unscoped().Delete(&Test{Id: 3})
	//db.Delete(&Test{Id: 4})
	//newTest := new(Test)
	//result := db.Where(&Test{Id: 5}).First(newTest)
	//var tests []Test
	//db.Find(&tests)
	//affected := result.RowsAffected
	//fmt.Println(affected)
}

func MyCache(db int) (cache *redis.Client) {
	option := &redis.Options{
		Addr:     "192.168.50.16:6379",
		Password: "PW2UFP7hA",
		DB:       db,
	}
	cache = redis.NewClient(option)
	return
}

func init() {
	println("atore")
}

func MyAuto() {
	test := new(entity.Order)
	user := new(entity.User)
	if err := DB.AutoMigrate(test); err != nil {
		fmt.Println("error")
	}
	if err := DB.AutoMigrate(user); err != nil {
		fmt.Println("error")
	}
}
