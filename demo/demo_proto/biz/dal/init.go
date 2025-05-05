package dal

import (
	"github.com/yydstrive/cloudwegomall/demo/demo_proto/biz/dal/mysql"
	//"github.com/yydstrive/cloudwegomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
