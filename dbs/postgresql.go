package dbs

import (
	"database/sql"
	"fmt"
)

type PostgresqlClient struct {
	Client *sql.DB
}

var pgdb *PostgresqlClient

func NewPostgresql() *sql.DB {
	db, err := sql.Open("postgres", "")
	if err != nil {
		fmt.Println("Postgresql 数据库创建失败!")
		return nil
	}
	// 校验dsn是否正确
	err = db.Ping()
	if err != nil {
		fmt.Printf("Postgresql 数据库创建失败:%s", err)
		return nil
	}

	ordb.Client = db
	fmt.Println("Postgresql 数据库初始化连接成功!")
	return db
}

func (pgdb *PostgresqlClient) Close() {
	pgdb.Client.Close()
}
func (pgdb *PostgresqlClient) Query() {
	pgdb.Client.Close()
}
