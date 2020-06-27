package logic

import (
	"testing"
	"github.com/treeforest/renju-server/handler/registry/proto"
	"github.com/treeforest/gos/client"
	"github.com/golang/protobuf/proto"
	"github.com/treeforest/renju-server/server/serviceId"
	"time"
)

func TestRegistryAccount_GetSmsCode(t *testing.T) {
	c := client.NewClient()
	c.Dial("127.0.0.1:9999")
	req := &registry.GetSmsCodeRequest{}
	phoneNum := "17859930679"
	req.PhoneNum = &phoneNum
	data, err := proto.Marshal(req)
	if err != nil {
		t.Error("proto marshal error: ", err)
	}
	t.Log("Send begin...")
	c.Send(uint32(serviceId.ID_REGISTRY), uint32(registry.Event_GET_SMS_CODE), data)
	t.Log("Send end...")
	for {
		msg := c.Recv()
		if msg == nil {
			time.Sleep(time.Millisecond * 50)
			continue
		}
		t.Log("Recv end...")
		resp := &registry.GetSmsCodeResponse{}
		if err := proto.Unmarshal(msg.GetData(), resp); err != nil {
			t.Error("proto unmarshal error: ", err)
		}
		t.Logf("Result:%d ExpireTime:%d", resp.GetResult(), resp.GetExpireTime())
	}
}

func TestRegistryAccount_GetMailCode(t *testing.T) {
	c := client.NewClient()
	c.Dial("127.0.0.1:9999")
	req := &registry.GetMailCodeRequest{}
	*req.Email = "1451670604@qq.com"
	data, err := proto.Marshal(req)
	if err != nil {
		t.Error("proto marshal error: ", err)
	}
	c.Send(uint32(serviceId.ID_REGISTRY), uint32(registry.Event_GET_MAIL_CODE), data)
	for {
		msg := c.Recv()
		if msg == nil {
			time.Sleep(time.Millisecond * 50)
			continue
		}
		resp := &registry.GetMailCodeResponse{}
		if err := proto.Unmarshal(msg.GetData(), resp); err != nil {
			t.Error("proto unmarshal error: ", err)
		}
		t.Logf("Result:%d ExpireTime:%d", *resp.Result, *resp.ExpireTime)
	}
}