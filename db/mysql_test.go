package db

import "testing"

func TestNewDB(t *testing.T) {
	db := NewDB("mysql", "root:@tcp(127.0.0.1:4000)/game?charset=utf8", "Asia/Shanghai")
	if err := db.Ping(); err != nil {
		t.Error(err)
	}
}
