package Mapping

import "phoenixbuilder/fastbuilder/sync_map"

// 指代 []byte 的 hex 形式
type HexString string

// 描述一个以 HexString 为键的 Map 。
// 其值将始终使用空结构体 struct{}
type Mapping struct {
	contents sync_map.Map[HexString, struct{}]
}