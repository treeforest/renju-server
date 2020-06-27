package logic

import (
	"github.com/treeforest/renju-server/handler/registry/proto"
	"github.com/treeforest/renju-server/global"
	"github.com/treeforest/logger"
	"time"
	"encoding/json"
	"fmt"
)

var (
	mgDatabase = "registry"
	mgIndexTable = "index_id" // 自增序号表
	mgPhoneCodeTable = "phone_code" // 手机验证码缓存表
	mhEmailCodeTable = "email_code" // 邮箱验证码缓存表
)

// redis-存储手机验证码过期信息
type phoneNumberCodeInfo struct {
	Code       string `json:"Code"` // 验证码
	ExpireTime int64  `json:"expireTime"`   // 过期时间
}

// redis-存储邮箱验证码过期信息
type emailCodeInfo struct {
	Code       string `json:"Code"` // 验证码
	ExpireTime int64  `json:"expireTime"`   // 过期时间
}

// mongodb-自增序号结构
type indexTable struct {
	ID   uint64 `bson:"_id"`
	Name string `bson:"Name"`
}

func GetIncID(Name string) uint64 {
	return GetIncIDBase(Name, 0)
}

func GetIncIDBase(Name string, baseId uint64) uint64 {
	//index := indexTable{}
	//err := global.Mongodb().FindAndInc(mgDatabase, )
	return 0
}

// 注册账号
type RegistryAccount struct {}

// 注册账号
func (r *RegistryAccount) Registry(req *registry.RegisterRequest) (resp *registry.RegisterResponse) {
	resp = new(registry.RegisterResponse)
	now := time.Now().Unix()
	var result int32
	defer func() {
		resp.Result = &result
	}()

	// 1、从redis中获取数据
	key := getRedisKey(req.GetAccount())
	val, err := global.Redis().GetKey(key)
	if err != nil {
		log.Debugf("GetKey %s error: %v", req.GetAccount(), err)
		result = int32(registry.Code_ERR_SVR_FATAL)
		return
	}

	// 2、解析出过期时间和已发送的验证码
	var code string
	var expireTime int64
	switch req.GetAccountType() {
	case registry.AccountType_PHONE:
		info := phoneNumberCodeInfo{}
		if err = json.Unmarshal([]byte(val), &info); err != nil {
			log.Debugf("Json Unmarshal error: %v", err)
			result = int32(registry.Code_ERR_UNMARSHAL)
			return
		}
		expireTime = info.ExpireTime
		code = info.Code
	case registry.AccountType_EMAIL:
		info := emailCodeInfo{}
		if err = json.Unmarshal([]byte(val), &info); err != nil {
			log.Debugf("Json Unmarshal error: %v", err)
			result = int32(registry.Code_ERR_UNMARSHAL)
			return
		}
		expireTime = info.ExpireTime
		code = info.Code
	case registry.AccountType_VISITOR:
		// 游客直接返回一个ID
		break
	}

	// 3、判断验证码是否匹配
	if code != req.GetCode() {
		log.Debugf("Code verify failed")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	// 4、判断验证码是否过期
	if now > expireTime {
		log.Debugf("Code already timeout")
		result = int32(registry.Code_ERR_CODE_OUTTIME)
		return
	}

	// 5、匹配成功
	// TODO: 存入用户信息服

	return
}

// 获取手机验证码
func (r *RegistryAccount) GetSmsCode(req *registry.GetSmsCodeRequest) (resp *registry.GetSmsCodeResponse) {
	resp = new(registry.GetSmsCodeResponse)
	var err error
	var ok bool
	var result int32
	defer func() {
		resp.Result = &result
	}()

	log.Debugf("GetSmsCode Begin: phoneNumber[%s]", req.GetPhoneNum())

	// 1、检查电话号码格式
	if ok = VerifyMobileFormat(req.GetPhoneNum()); !ok {
		log.Debugf("VerifyMobileFormat error: %s", req.GetPhoneNum())
		result = int32(registry.Code_ERR_PHONENUM_FORMAT)
		return
	}

	// 2、判断该电话号码是否已在redis，且还未过期
	info := &phoneNumberCodeInfo{}
	val, _ := global.Redis().GetKey(getRedisKey(req.GetPhoneNum()))
	if val != "" {
		if err = json.Unmarshal([]byte(val), info); err != nil {
			log.Debugf("Json Unmarshal error: %v", err)
			result = int32(registry.Code_ERR_UNMARSHAL)
			return
		}

		if info.ExpireTime > time.Now().Unix() {
			// 验证码未过期
			log.Debugf("PhoneNumber[%s] Code has not expired", req.GetPhoneNum())
			result = int32(registry.Code_ERR_CODE_UNEXPIRE)
			return
		}
	} else {
		// TODO: 判断是否已注册

	}

	// TODO：判断该号码是否被加入了黑名单

	// 3、生成验证码
	code := global.GenerateCode()

	// 4、将验证码发送给阿里云短信服务端（测试暂不使用）
	//if ok := global.SendSms(req.GetPhoneNum(), Code); !ok {
	//	log.Error("SendSms error")
	//	result = int32(registry.Code_ERR_SVR_FATAL)
	//	return
	//}
	log.Debugf("GetSmsCode phoneNumber:%s Code:%s", req.GetPhoneNum(), code)

	// 5、将验证码和信息存储到redis中
	var expireTime int64 = 60 * 5 // 过期时间为五分钟
	info.Code = code
	info.ExpireTime = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	b, _ := json.Marshal(info)
	if err = global.Redis().SetKey(getRedisKey(req.GetPhoneNum()), string(b)); err != nil {
		log.Errorf("SetKey error: %v", err)
		result = int32(registry.Code_ERR_REDIS_FATAL)
		return
}
	log.Debugf("SetKey:%s Value:%s", getRedisKey(req.GetPhoneNum()), string(b))

	// 6、操作成功
	result = int32(registry.Code_SUCCESS)
	resp.ExpireTime = &expireTime

	return
}

// 验证手机验证码
func (r *RegistryAccount) PreCheckCode(req *registry.PreCheckSmsCodeRequest) (resp *registry.PreCheckSmsCodeResponse) {
	resp = new(registry.PreCheckSmsCodeResponse)
	var result int32
	defer func() {
		resp.Result = &result
	}()

	// 1、检测手机号码格式
	if ok := VerifyMobileFormat(req.GetPhoneNum()); !ok {
		log.Debugf("VerifyMobileFormat error: %s", req.GetPhoneNum())
		result = int32(registry.Code_ERR_PHONENUM_FORMAT)
		return
	}

	// 2、从redis获取记录的数据信息
	val, _ := global.Redis().GetKey(getRedisKey(req.GetPhoneNum()))
	if val == "" {
		log.Debug("redis not exist key")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	// 3、反序列化到结构体
	info := phoneNumberCodeInfo{}
	if err := json.Unmarshal([]byte(val), &info); err != nil {
		log.Debugf("Json Unmarshal error: %v", err)
		result = int32(registry.Code_ERR_UNMARSHAL)
		return
	}

	// 4、是否过期
	if time.Now().Unix() > info.ExpireTime {
		log.Debug("Code expired")
		result = int32(registry.Code_ERR_CODE_EXPIRE)
		return
	}

	// 5、验证码是否匹配
	if req.GetSmsCode() != info.Code {
		log.Debug("Sms Code verify failed")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	result = int32(registry.Code_SUCCESS)

	return
}

// 获取邮箱验证码
func (r *RegistryAccount) GetMailCode(req *registry.GetMailCodeRequest) (resp *registry.GetMailCodeResponse) {
	resp = new(registry.GetMailCodeResponse)
	var result int32
	defer func() {
		resp.Result = &result
	}()

	// 1、检测 email 格式是否准确
	if ok := VerifyEmailFormat(req.GetEmail()); !ok {
		log.Debugf("VerifyEmailFormat error")
		result = int32(registry.Code_ERR_EMAIL_FORMAT)
		return
	}

	// 2、是否已经发送验证码
	info := emailCodeInfo{}
	val, _ := global.Redis().GetKey(getRedisKey(req.GetEmail()))
	if val != "" {
		if err := json.Unmarshal([]byte(val), &info); err != nil {
			log.Debugf("Json Unmarshal error: %v", err)
			result = int32(registry.Code_ERR_UNMARSHAL)
			return
		}

		if info.ExpireTime > time.Now().Unix() {
			// 验证码未过期
			log.Debug("Phone number Code has not expired")
			result = int32(registry.Code_ERR_CODE_UNEXPIRE)
			return
		}
	}

	// TODO：判断该邮箱是否被加入了黑名单

	// 3、生成验证码
	code := global.GenerateCode()

	// 4、将验证码发送给用户
	content := fmt.Sprintf("[%s]尊敬的用户：您的验证码为 %s，该验证码5分钟内有效，请勿泄露他人！", global.AppName, code)
	if err := global.SendEmail([]string{req.GetEmail()}, content); err != nil {
		log.Errorf("SendEmail error: %v", err)
		result = int32(registry.Code_ERR_SVR_FATAL)
		return
	}

	log.Debugf("GetMailCode email:%s Code:%s", req.GetEmail(), code)

	// 5、将验证码和信息存储到redis中
	var expireTime int64 = 60 * 5 // 过期时间为五分钟
	info.Code = code
	info.ExpireTime = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	b, _ := json.Marshal(info)
	if err := global.Redis().SetKey(getRedisKey(req.GetEmail()), string(b)); err != nil {
		log.Errorf("SetKey error: %v", err)
		result = int32(registry.Code_ERR_REDIS_FATAL)
		return
	}

	// 6、操作成功
	result = int32(registry.Code_SUCCESS)
	resp.ExpireTime = &expireTime

	return
}

//验证邮箱验证码
func (r *RegistryAccount) PreCheckMailCode(req *registry.PreCheckCMailCodeRequest) (resp *registry.PreCheckCMailCodeResponse) {
	resp = new(registry.PreCheckCMailCodeResponse)
	var result int32
	defer func() {
		resp.Result = &result
	}()

	// 1、检测邮箱格式
	if ok := VerifyEmailFormat(req.GetEmail()); !ok {
		log.Debugf("VerifyEmailFormat error: %s", req.GetEmail())
		result = int32(registry.Code_ERR_EMAIL_FORMAT)
		return
	}

	// 2、从redis获取记录的数据信息
	val, _ := global.Redis().GetKey(getRedisKey(req.GetEmail()))
	if val == "" {
		log.Debug("redis not exist key")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	// 3、反序列化到结构体
	info := emailCodeInfo{}
	if err := json.Unmarshal([]byte(val), &info); err != nil {
		log.Debugf("Json Unmarshal error: %v", err)
		result = int32(registry.Code_ERR_UNMARSHAL)
		return
	}

	// 4、是否过期
	if time.Now().Unix() > info.ExpireTime {
		log.Debug("Code expired")
		result = int32(registry.Code_ERR_CODE_EXPIRE)
		return
	}

	// 5、验证码是否匹配
	if req.GetEmailCode() != info.Code {
		log.Debug("Sms Code verify failed")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	result = int32(registry.Code_SUCCESS)

	return
}