package server

import (
	"github.com/treeforest/gos/transport"
	)

var globalServer transport.Server

/*
 * 获取服务实例
 */
func Server() transport.Server {
	return globalServer
}

/*
 * 启动服务
 */
func ServerStart(srvName string) {
	globalServer = transport.NewServer(srvName)

	globalServer.SetOnConnStartFunc(onConnStartFunc)
	globalServer.SetOnConnStopFunc(onConnStopFunc)

	registerRouter()

	globalServer.Serve()
}