package dal

import (
	"github.com/SGNYYYY/gomall/app/checkout/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
