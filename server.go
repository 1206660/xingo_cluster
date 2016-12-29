package main

import (
	"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/sys_rpc"
	"os"
	"path/filepath"
	"xingo_cluster/net_server"
	"xingo_cluster/gate_server"
	"xingo_cluster/admin_server"
        _ "net/http"
	_ "net/http/pprof"
)

func main() {
	//pprof
	//go func() {
	//	println(http.ListenAndServe("localhost:6060", nil))
	//}()

	//server code
	args := os.Args
	dir, err := filepath.Abs(filepath.Dir("."))
	if err == nil{
		s := clusterserver.NewClusterServer(args[1], filepath.Join(dir, "conf", "clusterconf.json"))
		s.AddRpcRouter(&sys_rpc.ChildRpc{})
		s.AddRpcRouter(&sys_rpc.RootRpc{})
		/*
		注册分布式服务器
		*/
		//net server
		s.AddModule("net", &net_server.TestNetApi{}, &net_server.TestNetRpc{})
		//gate server
		s.AddModule("gate", nil, &gate_server.TestGateRpc{})
		//admin server
		s.AddModule("admin", &admin_server.TestAdminHttp{}, nil)

		s.StartClusterServer()
	}
}