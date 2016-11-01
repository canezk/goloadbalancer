package zookeeper

import (
	"testing"
)

func TestConnect(t *testing.T) {
	var hosts []string = []string{"localhost:2181"}
	_, err := Connect(hosts)
	if err != nil {
		t.Errorf("Connect failed %s", err)
	}
}
