package dal

import (
	"github.com/SGNYYYY/gomall/app/cart/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
