package mysql

import (
	"github.com/yydstrive/cloudwegomall/demo/demo_proto/conf"
	"github.com/yydstrive/cloudwegomall/demo/demo_proto/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	type Version struct {
		Version string
	}
	var v Version
	err = DB.Raw("select version() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	fmt.Println(v)
}
