package registry_route

import (
	"github.com/treeforest/gos/transport"
	"github.com/treeforest/renju-server/handler/registry/logic"
	"github.com/treeforest/renju-server/handler/registry/proto"
	"github.com/golang/protobuf/proto"
	"github.com/treeforest/logger"
)

// 初始化一个处理业务的实例对象
var registry_handle = logic.RegistryAccount{}

type RegistryRoute struct {
	transport.BaseRouter
}

// 处理业务之前的方法
func (r *RegistryRoute) PreHandle(req transport.Request) {
	// TODO:
}

/*
 * 处理业务的方法
 *
 * 主要是对 req.GetData() 反序列化, 使用自己内部的 ID
 * 然后根据ID进行分发数据
 */
func (r *RegistryRoute) Handle(req transport.Request) {
	var data []byte
	var err error = nil

	// 根据方法ID进行路由
	switch registry.Event(req.GetMethodID()) {
	case registry.Event_REGISTRY:
		r := new(registry.RegisterRequest)
		var resp *registry.RegisterResponse
		if err = proto.Unmarshal(req.GetData(), r); err != nil {
			log.Errorf("registry.RegisterRequest Unmarshal error: %v", err)
			resp = new(registry.RegisterResponse)
			*resp.Result = int32(registry.Code_ERR_UNMARSHAL)
		} else {
			resp = registry_handle.Registry(r)
		}

		data, err = proto.Marshal(resp)
		if err != nil {
			log.Errorf("registry.RegisterResponse Marshal error: %v", err)
		}

	case registry.Event_GET_SMS_CODE:
		r := new(registry.GetSmsCodeRequest)
		var resp *registry.GetSmsCodeResponse
		if err = proto.Unmarshal(req.GetData(), r); err != nil {
			log.Errorf("registry.GetSmsCodeRequest Unmarshal error: %v", err)
			resp = new(registry.GetSmsCodeResponse)
			*resp.Result = int32(registry.Code_ERR_UNMARSHAL)
		} else {
			resp = registry_handle.GetSmsCode(r)
		}

		data, err = proto.Marshal(resp)
		if err != nil {
			log.Errorf("registry.GetSmsCodeResponse Marshal error: %v", err)
		}

	case registry.Event_PRE_CHECK_CODE:
		r := new(registry.PreCheckSmsCodeRequest)
		var resp *registry.PreCheckSmsCodeResponse
		if err = proto.Unmarshal(req.GetData(), r); err != nil {
			log.Errorf("registry.PreCheckSmsCodeRequest Unmarshal error: %v", err)
			resp = new(registry.PreCheckSmsCodeResponse)
			*resp.Result = int32(registry.Code_ERR_UNMARSHAL)
		} else {
			resp = registry_handle.PreCheckCode(r)
		}

		data, err = proto.Marshal(resp)
		if err != nil {
			log.Errorf("registry.PreCheckSmsCodeResponse Marshal error: %v", err)
		}

	case registry.Event_GET_MAIL_CODE:
		r := new(registry.GetMailCodeRequest)
		var resp *registry.GetMailCodeResponse
		if err = proto.Unmarshal(req.GetData(), r); err != nil {
			log.Errorf("registry.GetMailCodeRequest Unmarshal error: %v", err)
			resp = new(registry.GetMailCodeResponse)
			*resp.Result = int32(registry.Code_ERR_UNMARSHAL)
		} else {
			resp = registry_handle.GetMailCode(r)
		}

		data, err = proto.Marshal(resp)
		if err != nil {
			log.Errorf("registry.GetMailCodeResponse Marshal error: %v", err)
		}

	case registry.Event_PRE_CHECK_MAIL_CODE:
		r := new(registry.PreCheckCMailCodeRequest)
		var resp *registry.PreCheckCMailCodeResponse
		if err = proto.Unmarshal(req.GetData(), r); err != nil {
			log.Errorf("registry.PreCheckCMailCodeRequest Unmarshal error: %v", err)
			resp = new(registry.PreCheckCMailCodeResponse)
			*resp.Result = int32(registry.Code_ERR_UNMARSHAL)
		} else {
			resp = registry_handle.PreCheckMailCode(r)
		}

		data, err = proto.Marshal(resp)
		if err != nil {
			log.Errorf("registry.PreCheckCMailCodeResponse Marshal error: %v", err)
		}
	}

	if err = req.GetConnection().Send(req.GetServiceID(), req.GetMethodID(), data); err != nil {
		log.Warnf("Send error: %v", err)
	}
}

// 处理业务后的方法
func (r *RegistryRoute) PostHandle(req transport.Request) {
	// TODO:
}
