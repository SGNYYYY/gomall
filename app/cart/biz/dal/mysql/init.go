package mysql

import (
	"fmt"
	"os"

	"github.com/SGNYYYY/gomall/app/cart/biz/model"
	"github.com/SGNYYYY/gomall/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Cart{})
		DB.AutoMigrate( //nolint:errcheck
			&model.Cart{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `cart`.`cart` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','1','1','1'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','1','2','1')")
		}
	}
}
