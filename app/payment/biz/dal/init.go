package dal

import (
	"github.com/SGNYYYY/gomall/app/payment/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
