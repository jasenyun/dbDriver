package dbs

import (
	"database/sql"
	"fmt"
)

type MysqlClient struct {
	Client *sql.DB
}

var sqlDb *MysqlClient

func NewMySql() *sql.DB {
	db, err := sql.Open("mysql", "")
	if err != nil {
		fmt.Println("MySQL 数据库创建失败!")
		return nil
	}
	// 校验dsn是否正确
	err = db.Ping()
	if err != nil {
		fmt.Println("MySQL 数据库创建失败!")
		return nil
	}
	sqlDb.Client = db
	fmt.Println("MySQL 数据库初始化连接成功!")
	return db
}

func (sqlDb *MysqlClient) Close() {
	sqlDb.Client.Close()
}
func (sqlDb *MysqlClient) Query() {
	sqlDb.Client.Close()
}
