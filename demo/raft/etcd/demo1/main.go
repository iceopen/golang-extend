package main

import (
	"fmt"

	"go.etcd.io/etcd/raft"
)

func main() {
	storage := raft.NewMemoryStorage()
	// Recover the in-memory storage from persistent snapshot, state and entries.
	// 根据快照、entry日志等恢复当前raft节点到之前的状态
	storage.ApplySnapshot(snapshot)
	storage.SetHardState(state)
	storage.Append(entries)
	c := &raft.Config{
		//代表一个节点的ID，必须唯一，并且不能为0，不能重复利用，和zookeeper的id类似
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	//设置节点列表
	n := raft.StartNode(c, []raft.Peer{{ID: 0x02}, {ID: 0x03}})
	fmt.Println(n)
}
