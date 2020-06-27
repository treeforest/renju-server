package global

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/treeforest/logger"
	"math/rand"
)

/*
 * 发送短信验证码（对单用户发送）
 *
 * 参数
 *  phoneNumber: 用户手机号
 *  code: 验证码（暂定为6位）
 */
func sendSms(phoneNumber string, code string) (ok bool) {
	// 请求头书写
	request := requests.NewCommonRequest()
	request.Scheme = "https"
	request.Method = "POST"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = phoneNumber                // 手机号
	request.QueryParams["SignName"] = "五子连珠"                         // 项目名
	request.QueryParams["TemplateCode"] = "SMS_193524108"            //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + code + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := globalSmsClient.ProcessCommonRequest(request)
	if err != nil {
		log.Errorf("SendSms error: %v", err)
		return false
	}
	if !response.IsSuccess() {
		log.Error("SendSms failed.")
		return false
	}
	return true
}

/*
 * 生成6位数随机验证码
 */
func generateSmsCode() string {
	r := rand.New(rand.NewSource(globalSnowflake.Generate()))
	return fmt.Sprintf("%06v", r.Int31n(1000000))
}
