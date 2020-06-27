package global

import (
	"github.com/treeforest/gos/utils/config"
	"github.com/treeforest/gos/utils/config/encoder/yaml"
	"github.com/treeforest/gos/utils/config/source"
	"github.com/treeforest/gos/utils/config/source/file"
	"github.com/treeforest/gos/utils/dao/cache/redis"
	"github.com/treeforest/gos/utils/dao/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path/filepath"
	"testing"
)

func TestSnowflake(t *testing.T) {
	_sf, _ := NewSnowflake(1)
	for i := 0; i < 10; i++ {
		t.Log(_sf.Generate())
	}
}

// 测试数据库/缓存组件
func TestDao(t *testing.T) {
	type doc struct {
		Name string `bson:"Name" json:"Name"`
		Age  int    `bson:"Age" json:"Age"`
	}
	database := "test"
	table := "test_table"
	docs := doc{Name: "Tony", Age: 18}

	r := redis.NewRedisObj("127.0.0.1:6379", "", 128)

	//b, _ := json.Marshal(docs)
	//r.LPush("test_key", string(b))

	if err := r.LPush("test_key", "test_value"); err != nil {
		t.Errorf("redis LPush error: %v", err)
	} else {
		r.LPop("test_key")
	}

	m := mongodb.NewMongoCenter([]string{"127.0.0.1", "27017"}, "", "", 1024)
	if err := m.Insert(database, table, docs); err != nil {
		t.Errorf("mongodb insert error: %v", err)
	} else {
		if err = m.RemoveAll(database, table, bson.M{"Name": docs.Name}); err != nil {
			t.Errorf("mongodb remove error: %v", err)
		}
	}
}

// 测试配置文件的读写
func TestConfig(t *testing.T) {
	// 创建配置文件对象
	conf, err := config.NewConfig()
	if err != nil {
		t.Errorf("new config error: %v", err)
	}

	// 加载配置文件
	path := filepath.Join(os.Getenv("GoPath"), "src", "github.com", "treeforest", "renju-server", "config", "renju.yaml")
	if err = conf.Load(file.NewSource(
		file.WithPath(path),
		source.WithEncoder(yaml.NewEncoder()),
	)); err != nil {
		t.Errorf("config load error: %v", err)
	}

	// 读取配置文件
	redisAddr := conf.Get("RedisAddr").String("0.0.0.0:6379")
	redisPassword := conf.Get("RedisPassword").String("")
	mongodbAddr := conf.Get("MongodbAddr").String("")
	mongodbUser := conf.Get("MongodbUser").String("")
	mongodbPassword := conf.Get("MongodbPassword").String("")
	mongodbPoolLimit := conf.Get("MongodbPoolLimit").Int(0)
	t.Log(redisAddr, redisPassword, mongodbAddr, mongodbUser, mongodbPassword, mongodbPoolLimit)
}

// 测试邮箱发送
func TestSendEmail(t *testing.T) {
	content := "[五子连珠]尊敬的用户：您的验证码为 134213，该验证码 5 分钟内有效，请勿泄露他人。"
	if err := SendEmail([]string{"1451670604@qq.com", "384829308@qq.com"}, content); err != nil {
		t.Error(err)
	}
}
