package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

type OracleClient struct {
	Client *sql.DB
}

var ordb *OracleClient

func NewOracle() *sql.DB {
	db, err := sql.Open("godror", "")
	if err != nil {
		fmt.Println("ORACLE 数据库创建失败!")
		return nil
	}
	// 校验dsn是否正确
	err = db.Ping()
	if err != nil {
		fmt.Printf("ORACLE 数据库创建失败:%s", err)
		return nil
	}

	ordb.Client = db
	fmt.Println("ORACLE 数据库初始化连接成功!")
	return db
}

func (ordb *OracleClient) Close() {
	ordb.Client.Close()
}
func (ordb *OracleClient) Query() {
	sqlDb.Client.Close()
}
