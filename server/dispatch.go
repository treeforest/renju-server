package server

import (
	"github.com/treeforest/renju-server/server/serviceId"
	"github.com/treeforest/renju-server/handler/registry/route"
)

/*
 * 注册服务
 *
 * 参数
 *  serviceID: 服务唯一ID
 *  router: 服务实现的内部路由
 */
func registerRouter() {
	// 注册服
	Server().RegisterRouter(uint32(serviceId.ID_REGISTRY), &registry_route.RegistryRoute{})
}
