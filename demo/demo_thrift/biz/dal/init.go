package dal

import (
	"github.com/yydstrive/cloudwegomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/yydstrive/cloudwegomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
