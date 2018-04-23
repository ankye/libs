package db

import (
	"fmt"
	"time"
	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//NewDB create MYSQL xorm engine
func NewDB(driver string, dsm string, timezone string) *xorm.Engine {

	engine, err := xorm.NewEngine(driver, dsm)
	//engine.ShowSQL(true)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err.Error())
	}
	engine.TZLocation, _ = time.LoadLocation(timezone)
	return engine
}
