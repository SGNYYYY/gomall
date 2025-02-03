package dal

import (
	"github.com/SGNYYYY/gomall/app/order/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
