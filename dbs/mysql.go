package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlClient struct {
	Client *sql.DB
}

var sqlDb *MysqlClient

func NewMySql(dsn string) *MysqlClient {
	db, err := sql.Open("mysql", dsn)
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
	sqlDb = &MysqlClient{}
	sqlDb.Client = db
	fmt.Println("MySQL 数据库初始化连接成功!")
	return sqlDb
}

func (sqlDb *MysqlClient) Close() {
	sqlDb.Client.Close()
}
func (sqlDb *MysqlClient) Query(sqlStr string) []map[string]string {
	res, err := sqlDb.Client.Query(sqlStr)
	if err != nil {
		return nil
	}
	columns, _ := res.Columns()
	count := len(columns)
	var values = make([]interface{}, count)
	var scanValuse = make([]interface{}, count)
	for i := range values {
		scanValuse[i] = &values[i]
	}
	i := 0
	record := make([]map[string]string, 0)
	for res.Next() {
		res.Scan(values)
		row := make(map[string]string)
		for i, v := range values {
			if v != nil {
				key := columns[i]
				row[key] = string(v.([]byte))
			}
		}
		record[i] = row
		i++
	}
	return record
}

func (sqlDb *MysqlClient) InsertOrUpdate(sqlStr string) int64 {
	result, err := sqlDb.Client.Exec(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	return rows
}

func (sqlDb *MysqlClient) Delete(sqlStr string) bool {
	result, _ := sqlDb.Client.Exec(sqlStr)
	rows, _ := result.RowsAffected()
	if rows > 0 {
		return true
	} else {
		return false
	}
}
