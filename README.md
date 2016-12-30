# xingo_cluster
xingo cluster 分布式集群 示例代码<br>
示例配置:<br>
```json
{
    "master":{"host": "192.168.2.225","rootport":9999},
    "servers":{
        "gate2":{"host": "192.168.2.225", "rootport":10000,"name":"gate2", "module": "gate", "log": "gate2.log"},
        "gate1":{"host": "192.168.2.225", "rootport":10001,"name":"gate1", "module": "gate", "log": "gate1.log"},
        "net1":{"host": "192.168.2.225", "netport":11009,"name":"net1","remotes":["gate2", "gate1"], 
                    "module": "net", "log": "net.log"},
        "net2":{"host": "192.168.2.225", "netport":11009,"name":"net2","remotes":["gate2", "gate1"], 
                    "module": "net", "log": "net.log"},
        "net3":{"host": "192.168.2.225", "netport":11009,"name":"net3","remotes":["gate2", "gate1"], 
                    "module": "net", "log": "net.log"},
        "net4":{"host": "192.168.2.225", "netport":11009,"name":"net4","remotes":["gate2", "gate1"], 
                    "module": "net", "log": "net.log"},
        "admin":{"host": "192.168.2.225", "remotes":["gate2", "gate1"], "name":"admin", "module": "admin", 
            "http": [8888, "/static"]},
        "game1":{"host": "192.168.2.225", "remotes":["gate2", "gate1"], "name":"game1", "module": "game"}
    }
}
```
架构图：<br>
![alt text](conf/xingo_cluster_架构.png)
