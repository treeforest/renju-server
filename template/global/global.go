package global

import (
	"github.com/treeforest/gos/utils/dao/cache/redis"
	"github.com/treeforest/gos/utils/dao/db/mongodb"
	"github.com/treeforest/gos/utils/config"
	"github.com/treeforest/logger"
	"path/filepath"
	"github.com/treeforest/gos/utils/config/encoder/yaml"
	"github.com/treeforest/gos/utils/config/source/file"
	"github.com/treeforest/gos/utils/config/source"
	"strings"
	"os"
)

/*
 * 全局函数变量
 */
var (
	defaultRedis redis.RedisOp
	defaultMongo mongodb.MongoOp
	defaultConfig config.Config
)

// Redis
func Redis() redis.RedisOp {
	return defaultRedis
}

// MongoDB
func Mongo() mongodb.MongoOp {
	return defaultMongo
}

// 初始化
func init() {
	// Init()
}

func Init() {
	if (!initConfig()) {
		return
	}

	redisAddr := defaultConfig.Get("RedisAddr").String("")
	redisPassword := defaultConfig.Get("RedisPassword").String("")
	if len(redisAddr) != 0 {
		defaultRedis = redis.NewRedisObj(redisAddr, redisPassword, 128)
	}

	mongodbAddr := defaultConfig.Get("MongodbAddr").String("")
	mongodbUser := defaultConfig.Get("MongodbUser").String("")
	mongodbPassword := defaultConfig.Get("MongodbPassword").String("")
	mongodbPoolLimit := defaultConfig.Get("MongodbPoolLimit").Int(1024)
	if len(mongodbAddr) != 0 {
		ms := strings.Split(mongodbAddr, ",")
		defaultMongo = mongodb.NewMongoCenter(ms, mongodbUser, mongodbPassword, mongodbPoolLimit)
	}
}

func initConfig() (ok bool){
	defaultConfig, err := config.NewConfig()
	if err != nil {
		log.Error("new config error: ", err)
		return false
	}

	// 配置文件路径
	path := filepath.Join(os.Getenv("GoPath"), "src", "github.com", "treeforest", "renju-server", "template", "config.yaml")

	if err = defaultConfig.Load(file.NewSource(
		file.WithPath(path),
		source.WithEncoder(yaml.NewEncoder()),
	)); err != nil {
		log.Errorf("config load error: %v", err)
	}

	return true
}