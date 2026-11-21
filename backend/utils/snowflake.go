package utils

import (
	"github.com/bwmarrin/snowflake"
)

// GenerateId 生成分布式id
// node 机器节点 [1:"id",2:"token_id",3:"uri"]
func GenerateId(node int64) int64 {
	n, _ := snowflake.NewNode(node)
	s := n.Generate()
	return int64(s)
}

// 生成特定前缀分布式字符串
func GenerateUri(pre string) string {
	n, _ := snowflake.NewNode(42)
	return pre + n.Generate().String()
}
