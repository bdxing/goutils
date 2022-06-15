package goutils

import "testing"

func TestStrHumpToUnder(t *testing.T) {
	t.Log(StrHumpToUnder("UserName: alan"))
}

func TestStrUnderToHump(t *testing.T) {
	t.Log(StrUnderToHump("user_name: alan"))
}
