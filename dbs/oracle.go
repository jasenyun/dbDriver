package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"log"
)

type OracleClient struct {
	Client *sql.DB
}

var ordb *OracleClient

func NewOracle(dsn string) *OracleClient {
	db, err := sql.Open("godror", dsn)
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
	ordb = &OracleClient{}
	ordb.Client = db
	fmt.Println("ORACLE 数据库初始化连接成功!")
	return ordb
}

func (ordb *OracleClient) Close() {
	ordb.Client.Close()
}
func (ordb *OracleClient) Query(sqlStr string) []map[string]string {
	res, err := ordb.Client.Query(sqlStr)
	if err != nil {
		return nil
	}
	columns, _ := res.Columns()
	count := len(columns)
	var values = make([]interface{}, count)
	var scanValuse = make([]interface{}, count)
	for i, _ := range values {
		scanValuse[i] = &values[i]
	}
	i := 0
	record := make([]map[string]string, 0)
	for res.Next() {
		res.Scan(scanValuse)
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

func (ordb *OracleClient) InsertOrUpdate(sqlStr string) int64 {
	result, err := ordb.Client.Exec(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	return rows
}

func (ordb *OracleClient) Delete(sqlStr string) bool {
	result, _ := ordb.Client.Exec(sqlStr)
	rows, _ := result.RowsAffected()
	if rows > 0 {
		return true
	} else {
		return false
	}
}
