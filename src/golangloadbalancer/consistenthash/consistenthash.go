/**
一致性哈希算法
 */
package consistenthash

import (
	"log"
	"github.com/serialx/hashring"
	"golangloadbalancer/zookeeper"
	"github.com/samuel/go-zookeeper/zk"
)

type Consistent struct{
	domainZnode map[string]string //域名和znode的映射关系
	domainRing map[string]hashring.HashRing //domain及其对应的哈希环
	conn *zk.Conn //zk的连接
	ZookeeperUrl []string //zookeeper的地址
}

func (c *Consistent)AddDomain(domain string, znode string) {
	c.domainZnode[domain] = znode
	log.Printf("Add %s znode and %s domain success!\n", domain, znode)
}

func (c *Consistent)AddDomainServers(domain string, znode string)  {
	state := zk.State(c.conn)
	for state != zk.StateConnected {
		c.conn, _ = zookeeper.Connect(c.ZookeeperUrl)
		state = zk.State(c.conn)
	}
	servers, err := zookeeper.GetServerList(c.conn, znode)
	if err != nil {
		log.Print("Add server of domain %s failed,caused by %v+\n", domain, err)
		return
	}
	c.domainRing[domain] = hashring.New(servers)
}

func (c *Consistent)GetServer(domain string, key string) (server string) {
	ring := c.domainRing[domain]
	server,_ = ring.GetNode(key)
	return
}

