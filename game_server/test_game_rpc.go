package game_server

import (
	"github.com/viphxin/xingo/cluster"
	"github.com/viphxin/xingo/logger"
	_"github.com/viphxin/xingo/clusterserver"
	"github.com/viphxin/xingo/clusterserver"
	"fmt"
	"strings"
)

type TestGameRpc struct {

}

func (this *TestGameRpc)Proxy2Game(request *cluster.RpcRequest){
	//Json反序列化数字到interface{}类型的值中，默认解析为float64类型，在使用时要注意
	pid := int32((request.Rpcdata.Args[0]).(float64))
	message := (request.Rpcdata.Args[1]).(string)
	logger.Debug(pid, message)
	//onenet := clusterserver.GlobalClusterServer.ChildsMgr.GetRandomChild("net")
	//if onenet != nil{
	//	onenet.CallChildNotForResult("PushMsg2Client", pid, message)
	//}
	for _, child := range clusterserver.GlobalClusterServer.RemoteNodesMgr.GetChilds(){
		if strings.Contains(child.GetName(), "net"){
			child.CallChildNotForResult("PushMsg2Client", pid, message)
		}
	}
}

func (this *TestGameRpc)Add(request *cluster.RpcRequest) map[string]interface{}{
	//Json反序列化数字到interface{}类型的值中，默认解析为float64类型，在使用时要注意
	i := int32((request.Rpcdata.Args[0]).(float64))
	ii := int32((request.Rpcdata.Args[1]).(float64))

	logger.Debug(fmt.Sprintf("%d + %d = %d", i, ii, i + ii))
	return map[string]interface{}{
		"sum": i + ii,
	}
}