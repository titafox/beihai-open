package tests

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	s := time.Now().Format("20060102")
	t.Log(s)
}

func TestTime2(t *testing.T) {
	s := time.Unix(time.Now().Unix()+300, 0).Format("2006-01-02 15:04:05")
	t.Log(s)
}
