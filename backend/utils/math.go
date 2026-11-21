package utils

import (
	"bytes"
	"encoding/binary"
)

const (
	// EXCHANGE_RATE eth整数换算费率
	EXCHANGE_RATE = 100000000
)

func Eth2Integer(eth float64) uint {
	return uint(eth * float64(EXCHANGE_RATE))
}

func Integer2Eth(eth uint) float64 {
	money := float64(eth) / float64(EXCHANGE_RATE)
	return money
}

func IntToBytes(n int) []byte {
	data := int64(n)
	byteBuf := bytes.NewBuffer([]byte{})
	binary.Write(byteBuf, binary.BigEndian, data)
	return byteBuf.Bytes()
}

// Int64ToBytes 64位整数转为字节
func Int64ToBytes(data int64) []byte {
	byteBuf := bytes.NewBuffer([]byte{})
	binary.Write(byteBuf, binary.BigEndian, data)
	return byteBuf.Bytes()
}
