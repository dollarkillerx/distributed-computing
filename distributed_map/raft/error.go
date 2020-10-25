package raft

import "errors"

var (
	ErrCompacted = errors.New("requested index is unavailable due to compaction")
)
