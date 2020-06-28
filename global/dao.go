package global

import (
	"fmt"
	"github.com/treeforest/gos/utils/dao/cache/redis"
	"github.com/treeforest/gos/utils/dao/db/mongodb"
	"strings"
)

// 数据库初始化
func initDatabase() error {
	// 初始化 Redis
	redisAddr := Config().Get("RedisAddr").String("0.0.0.0:6379")
	redisPassword := Config().Get("RedisPassword").String("")
	if len(redisAddr) == 0 {
		return fmt.Errorf("Redis init failed.")
	}
	globalRedis = redis.NewRedisObj(redisAddr, redisPassword, 128)

	// 初始化 Mongodb
	mongodbAddr := globalConfig.Get("MongodbAddr").String("")
	mongodbUser := globalConfig.Get("MongodbUser").String("")
	mongodbPassword := globalConfig.Get("MongodbPassword").String("")
	mongodbPoolLimit := globalConfig.Get("MongodbPoolLimit").Int(1024)
	if len(mongodbAddr) == 0 {
		return fmt.Errorf("Mongodb init failed.")
	}
	ms := strings.Split(mongodbAddr, ",")
	globalMongodb = mongodb.NewMongoCenter(ms, mongodbUser, mongodbPassword, mongodbPoolLimit)

	return nil
}
