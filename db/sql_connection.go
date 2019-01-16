package db

import (
	"database/sql"

	"github.com/gonethopper/libs/logs"
	"github.com/jmoiron/sqlx"
)

//NewSQLConnection create SQLConnection
func NewSQLConnection(driver string, dsn string, timezone string) *SQLConnection {
	s := new(SQLConnection)
	s.Driver = driver
	s.DSN = dsn
	return s
}

//SQLConnection 连接实体类
type SQLConnection struct {
	db *sqlx.DB

	Driver string
	DSN    string
}

//Connect 连接
func (s *SQLConnection) Connect() error {
	var err error
	if s.db, err = sqlx.Connect(s.Driver, s.DSN); err != nil {
		panic(err.Error())
	}
	return s.Ping()
}

//Ping test connection
func (s *SQLConnection) Ping() error {
	// force a connection and test ping
	err := s.db.Ping()
	if err != nil {
		logs.Error("couldn't connect to database: %s %s", s.Driver, s.DSN)
		panic(err.Error())
	}
	return err
}

//Close close connection
func (s *SQLConnection) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

//IsErrNoRows 判断是否有数据
func (s *SQLConnection) IsErrNoRows(err error) bool {
	return sql.ErrNoRows == err
}

//Select select操作
func (s *SQLConnection) Select(dest interface{}, query string, args ...interface{}) error {
	return s.db.Select(dest, query, args...)
}

//Exec 执行sql语句
func (s *SQLConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := s.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//QueryRow 查询单条记录
func (s *SQLConnection) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return s.db.QueryRowx(query, args...)
}

//Query 查询记录集
func (s *SQLConnection) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return s.db.Queryx(query, args...)
}
