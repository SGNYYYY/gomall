package dal

import (
	"github.com/SGNYYYY/gomall/app/email/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
