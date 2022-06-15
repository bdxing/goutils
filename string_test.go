package goutils

import (
	"math/rand"
	"testing"
	"time"
)

func TestStrHumpToUnder(t *testing.T) {
	t.Log(StrHumpToUnder("UserName: alan"))
}

func TestStrUnderToHump(t *testing.T) {
	t.Log(StrUnderToHump("user_name: alan"))
}

func TestStrRandom(t *testing.T) {
	// 添加随机种子
	rand.Seed(time.Now().UnixNano())

	t.Log(StrRandom(16, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789./-_&=")))
}
