/**
处理和zookeeper相关的操作:
1、和zk建立连接
2、注册watcher
3、读取znode上面的数据
 */
package zookeeper

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const timeOut  = 60

func Connect(hosts []string) (conn *zk.Conn, err error)  {
	conn, _, err = zk.Connect(hosts, timeOut*time.Second)
	if err != nil {
		fmt.Println("Connected failed to zookeeper!@" + err)
	}
	return
}

func GetServerList(conn *zk.Conn, path string) (list []string, err error) {
	list, _, err = conn.Children(path)
	return
}
