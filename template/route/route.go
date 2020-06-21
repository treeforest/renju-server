package route

import (
	"github.com/treeforest/gos/transport"
	"github.com/treeforest/renju-server/template/handler"
)

// 初始化一个处理业务的实例对象
var handle handler.Test = handler.Test{}

type TemplateRoute struct {
	transport.BaseRouter
}

// 处理业务之前的方法
func (r *TemplateRoute) PreHandle(req transport.Request) {
	// ...
}

/*
 * 处理业务的方法
 *
 * 主要是对 req.GetData() 反序列化, 使用自己内部的 ID
 * 然后根据ID进行分发数据
 */
func (r *TemplateRoute) Handle(req transport.Request) {
	var methodID int

	switch methodID {
	case 1:
		// 反序列话数据
		// 调用对应服务
		handle.SayHi(req)// req 不一定是 transport.Request 类型
	case 2:
		//
		handle.SayHello(req)
	case 3:
		handle.Jump(req)
	//...
	}
}

// 处理业务后的方法
func (r *TemplateRoute) PostHandle(req transport.Request) {

}
