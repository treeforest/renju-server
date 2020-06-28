package logic

import (
	"github.com/treeforest/renju-server/global/database"
	"github.com/treeforest/renju-server/handler/registry/proto"
	"github.com/treeforest/renju-server/global"
	"github.com/treeforest/logger"
	"time"
	"encoding/json"
	"fmt"
)

// 注册账号对象
type RegistryAccount struct {}

// 注册账号
func (r *RegistryAccount) Registry(req *registry.RegisterRequest) (resp *registry.RegisterResponse) {
	resp = new(registry.RegisterResponse)
	var result int32
	defer func() {
		resp.Result = &result
	}()

	// 1、从redis中获取数据
	val, err := db.Registry().GetKey(req.GetAccount())
	if err != nil {
		log.Debugf("GetKey %s error: %v", req.GetAccount(), err)
		result = int32(registry.Code_ERR_SVR_FATAL)
		return
	}

	// 2、解析验证码
	info := db.CodeExpireInfo{}
	if err = json.Unmarshal([]byte(val), &info); err != nil {
		log.Debugf("Json Unmarshal error: %v", err)
		result = int32(registry.Code_ERR_UNMARSHAL)
		return
	}

	// 3、判断验证码是否匹配
	if info.Code != req.GetCode() {
		log.Debugf("Code verify failed")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	// 4、存入用户信息服
	userInfo := &db.UserInfo{}
	userInfo.UserID = db.User().GetUserInfoIncID()
	userInfo.UserName = req.GetUserName()
	userInfo.Password = req.GetPassword()
	switch req.GetAccountType() {
	case registry.AccountType_PHONE:
		userInfo.PhoneNumber = req.GetAccount()
	case registry.AccountType_EMAIL:
		userInfo.Email = req.GetAccount()
	}

	if err := db.User().InsertUserInfo(userInfo); err != nil {
		log.Errorf("Insert user info error: %v", err)
		result = int32(registry.Code_ERR_SVR_FATAL)
		return
	}

	//TODO: 5、注册成功，设置登录状态为true，返回一个session code


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
	info := &db.CodeExpireInfo{}
	val, _ := db.Registry().GetKey(req.GetPhoneNum())
	if val != "" {
		// 验证码未过期
		log.Debugf("Phone number[%s] Code has not expired", req.GetPhoneNum())
		result = int32(registry.Code_ERR_CODE_UNEXPIRE)
		return
	}

	// 3、判断是否已经注册(是否在mongodb中）
	if ok := db.User().ExistPhoneNumber(req.GetPhoneNum()); ok {
		log.Debugf("Phone number[%s] has registered", req.GetPhoneNum())
		result = int32(registry.Code_ERR_ACCT_EXISTS)
		return
	}

	// 4、判断该号码是否被加入了黑名单
	if ok := db.User().ExistBlackListPhoneNumber(req.GetPhoneNum()); ok {
		log.Debugf("Phone number[%s] is on the blacklist", req.GetPhoneNum())
		result = int32(registry.Code_ERR_BLACK_LIST)
		return
	}

	// 5、生成验证码
	code := global.GenerateCode()

	// 6、将验证码发送给阿里云短信服务端（测试暂不使用）
	//if ok := global.SendSms(req.GetPhoneNum(), code); !ok {
	//	log.Error("SendSms error")
	//	result = int32(registry.Code_ERR_SVR_FATAL)
	//	return
	//}
	log.Debugf("GetSmsCode phoneNumber:%s Code:%s", req.GetPhoneNum(), code)

	// 7、将验证码和信息存储到redis中
	var expireTime int64 = 60 * 5 // 过期时间为五分钟
	info.Code = code
	info.ExpireTime = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	b, _ := json.Marshal(info)
	if err = db.Registry().SetExpireKey(req.GetPhoneNum(), string(b), int(info.ExpireTime) + 2); err != nil {
		log.Errorf("SetExpireKey error: %v", err)
		result = int32(registry.Code_ERR_REDIS_FATAL)
		return
	}
	log.Debugf("SetKey:%s Value:%s", db.Registry().RedisKey(req.GetPhoneNum()), string(b))

	// 8、操作成功
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
	val, _ := db.Registry().GetKey(req.GetPhoneNum())
	if val == "" {
		log.Debug("redis not exist key")
		result = int32(registry.Code_ERR_CODE_EXPIRE)
		return
	}

	// 3、反序列化到结构体
	info := db.PhoneNumberCodeInfo{}
	if err := json.Unmarshal([]byte(val), &info); err != nil {
		log.Debugf("Json Unmarshal error: %v", err)
		result = int32(registry.Code_ERR_UNMARSHAL)
		return
	}

	// 4、验证码是否匹配
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

	// 2、检测验证码是否过期
	info := db.CodeExpireInfo{}
	val, _ := db.Registry().GetKey(req.GetEmail())
	if val != "" {
		// 验证码未过期
		log.Debug("Email code has not expired")
		result = int32(registry.Code_ERR_CODE_UNEXPIRE)
		return
	}

	// 3、邮箱是否已经注册
	if ok := db.User().ExistEmail(req.GetEmail()); ok {
		log.Debugf("Email[%s] has registered", req.GetEmail())
		result = int32(registry.Code_ERR_EMAIL_BIND_ANOTHER)
		return
	}

	// 4、判断该邮箱是否被加入了黑名单
	if ok := db.User().ExistBlackListEmail(req.GetEmail()); ok {
		log.Debugf("Email[%s] is on the blacklist", req.GetEmail())
		result = int32(registry.Code_ERR_BLACK_LIST)
		return
	}

	// 5、生成验证码
	code := global.GenerateCode()

	// 6、将验证码发送给用户
	content := fmt.Sprintf("【%s】尊敬的用户：您的验证码为 %s，该验证码5分钟内有效，请勿泄露他人！", global.AppName, code)
	if err := global.SendEmail([]string{req.GetEmail()}, content); err != nil {
		log.Errorf("SendEmail error: %v", err)
		result = int32(registry.Code_ERR_SVR_FATAL)
		return
	}

	log.Debugf("GetMailCode email:%s Code:%s", req.GetEmail(), code)

	// 7、将验证码和信息存储到redis中
	var expireTime int64 = 60 * 5 // 过期时间为五分钟
	info.Code = code
	info.ExpireTime = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	b, _ := json.Marshal(info)
	if err := db.Registry().SetExpireKey(req.GetEmail(), string(b), int(expireTime) + 2); err != nil {
		log.Errorf("SetKey error: %v", err)
		result = int32(registry.Code_ERR_REDIS_FATAL)
		return
	}

	// 8、操作成功
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
	val, _ := db.Registry().GetKey(req.GetEmail())
	if val == "" {
		log.Debug("redis not exist key")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	// 3、反序列化到结构体
	info := db.CodeExpireInfo{}
	if err := json.Unmarshal([]byte(val), &info); err != nil {
		log.Debugf("Json Unmarshal error: %v", err)
		result = int32(registry.Code_ERR_UNMARSHAL)
		return
	}

	// 4、验证码是否匹配
	if req.GetEmailCode() != info.Code {
		log.Debug("Sms Code verify failed")
		result = int32(registry.Code_ERR_CODE_VERIFY)
		return
	}

	result = int32(registry.Code_SUCCESS)

	return
}