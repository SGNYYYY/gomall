package dal

import (
	"github.com/SGNYYYY/gomall/app/auth/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
