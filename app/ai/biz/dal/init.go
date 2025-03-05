package dal

import (
	"github.com/SGNYYYY/gomall/app/ai/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/ai/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
