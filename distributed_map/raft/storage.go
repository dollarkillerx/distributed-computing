package raft

import (
	pb "cmp/distributed_map/raft/raftpb"
	local "github.com/dollarkillerx/stele/internal/server"
	"github.com/dollarkillerx/stele/sdk/stele"
	"github.com/golang/protobuf/proto"
	"log"
)

// 存储是一个接口，可以由应用程序实现。
// 从存储中检索日志条目。
//
// 如果任何存储方法返回一个错误，筏子实例将
// 无法运作并拒绝参加选举；；
//在这种情况下，应用程序负责清理和恢复。
type Storage interface {
	// InitialState 返回保存的 HardState 和 ConfState 信息。
	InitialState() (pb.HardState, pb.ConfState, error)
	// Entries返回范围[lo,hi)内的日志条目片断。
	// MaxSize限制了返回的日志条目的总大小，但也有可能会限制日志条目的大小。
	// Entries返回至少一个条目（如果有的话）。
	Entries(lo, hi uint64) ([]pb.Entry, error)
	// Term 返回条目 i 的术语，其范围必须为
	// [FirstIndex()-1，LastIndex()]。前面的条目术语
	//FirstIndex被保留下来，以便进行匹配，即使是在
	//该条目的其余部分可能无法使用。
	Term(i uint64) (uint64, error)
	// LastIndex 返回日志中最后一个条目的索引。
	LastIndex() (uint64, error)
	// FirstIndex 返回第一个日志条目的索引，该索引是
	// 可能通过 "条目 "提供（旧的条目已被纳入）。
	// 到最新的快照中；如果存储只包含虚拟条目，那么
	//第一个日志条目不可用).FirstIndex() (uint64, error).
	FirstIndex() (uint64, error)
	// Snapshot返回最近的快照。
	// 如果快照暂时不可用，它应该返回ErrSnapshotTemporarilyUnavailable。
	// 所以筏子状态机可以知道存储需要一些时间来准备
	// 快照，然后再调用Snapshot。
	Snapshot() (pb.Snapshot, error)
}

type SteleStorage struct {
	local *local.Local
}

func NewSteleStorage() *SteleStorage {
	local, err := stele.NewLocal("SteleStorage")
	if err != nil {
		log.Fatalln(err)
	}
	return &SteleStorage{
		local: local,
	}
}

func (s *SteleStorage) InitialState() (hs pb.HardState, cs pb.ConfState, err error) {
	hardState, err := s.local.Get([]byte("hard_state"))
	if err != nil {
		log.Println(err)
		return hs, cs, nil
	}
	if err := proto.Unmarshal(hardState, &hs); err != nil {
		log.Println(err)
		return hs, cs, nil
	}

	confState, err := s.local.Get([]byte("conf_state"))
	if err != nil {
		log.Println(err)
		return hs, cs, nil
	}
	if err := proto.Unmarshal(confState, &cs); err != nil {
		log.Println(err)
		return hs, cs, nil
	}
	return hs, cs, nil
}

func (s *SteleStorage) SetHardState(hs pb.HardState) error {
	ms, err := proto.Marshal(&hs)
	if err != nil {
		return err
	}

	return s.local.Set([]byte("hard_state"), ms, 0)
}

func (s *SteleStorage) Entries()
