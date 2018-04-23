package db

import "testing"

func TestNewDB(t *testing.T) {
	db := NewDB("mysql", "root:godankye@1234@/mysql?charset=utf8", "Asia/Shanghai")
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}
