package db

import "github.com/treeforest/renju-server/global"

// redis-存储验证码过期信息
type CodeExpireInfo struct {
	Code       string `json:"Code"` // 验证码
	ExpireTime int64  `json:"expireTime"`   // 过期时间
}

// registry redis(Redis)
type registry struct {}

func (r *registry) RedisKey(item string) string {
	return "registry_" + item
}

func (r *registry) GetKey(key string) (val string, err error){
	return global.Redis().GetKey(r.RedisKey(key))
}

func (r *registry) SetExpireKey(sKey, sValue string, lExpireTime int) error {
	return global.Redis().SetExpireKey(sKey, sValue, lExpireTime)
}