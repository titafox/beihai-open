package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// LETTERS 定义原始字符串
const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// RandSeq 获取一个随机用户唯一编号
func RandSeq(codeLen int) string {
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	// 随机从中获取
	rand.Seed(time.Now().UnixNano())
	len := len(LETTERS)

	for rawStrLen := len; codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(LETTERS[randNum])
	}
	return b.String()
}

func RandomString(n int) string {
	randStr := make([]byte, n/2)
	rand.Read(randStr)
	return fmt.Sprintf("%x", randStr)
}

// Md5Str 字符串md5
func Md5Str(str string) string {
	var b []byte = []byte(str)
	data := []byte(b)
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	return s
}
