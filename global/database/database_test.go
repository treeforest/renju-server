package db

import "testing"

// 测试user mongodb
func TestUser(t *testing.T) {
	// index id
	t.Log(User().GetIncIDBase("test", 10000))
}
