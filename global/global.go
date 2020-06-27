package global

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/treeforest/gos/utils/config"
	"github.com/treeforest/gos/utils/dao/cache/redis"
	"github.com/treeforest/gos/utils/dao/db/mongodb"
	"github.com/treeforest/logger"
)

// 全局参数定义
var (
	globalSnowflake *snowflake
	globalSmsClient *sdk.Client
	globalConfig    config.Config

	globalRedis   redis.RedisOp
	globalMongodb mongodb.MongoOp

	AppName string // 应用名
)

// 返回雪花算法实例
func SnowFlake() *snowflake {
	return globalSnowflake
}

// 返回配置文件实例
func Config() config.Config {
	return globalConfig
}

// 返回redis实例
func Redis() redis.RedisOp {
	return globalRedis
}

// 返回mongodb实例
func Mongodb() mongodb.MongoOp {
	return globalMongodb
}

// 发送短信验证码
func SendSms(phoneNumber string, code string) (ok bool) {
	return sendSms(phoneNumber, code)
}

// 发送邮箱验证码
func SendEmail(toEmail []string, content string) error {
	return sendEmail(toEmail, content)
}

// 生成短信验证码
func GenerateCode() string {
	return generateSmsCode()
}

// 初始化操作
func globalInit() error {
	var err error = nil

	// 初始化配置文件
	if err = initConfig(); err != nil {
		log.Errorf("initConfig error: %v", err)
		return err
	}

	// 初始化应用名
	AppName = Config().Get("AppName").String("五子连珠")

	// 初始化雪花算法
	if globalSnowflake, err = NewSnowflake(101); err != nil {
		log.Errorf("NewSnowflake error: %v", err)
		return err
	}

	// 初始化sms服务
	accessKeyId := Config().Get("AccessKeyId").String("")
	accessKeySecret := Config().Get("AccessKeySecret").String("")
	if globalSmsClient, err = sdk.NewClientWithAccessKey("cn-beijing", accessKeyId, accessKeySecret); err != nil {
		log.Errorf("NewClientWithAccessKey error: %v", err)
		return err
	}

	// 初始化数据库操作
	if err = initDatabase(); err != nil {
		log.Errorf("Init database error: %v", err)
		return err
	}

	return nil
}

func init() {
	if err := globalInit(); err != nil {
		panic(err)
	}
}
