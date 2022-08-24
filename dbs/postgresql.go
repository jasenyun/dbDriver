package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PostgresqlClient struct {
	Client *sql.DB
}

var pgdb *PostgresqlClient

func NewPostgresql(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
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

	fmt.Println("Postgresql 数据库初始化连接成功!")
	return db
}

func (pgdb *PostgresqlClient) Close() {
	pgdb.Client.Close()
}
func (pgdb *PostgresqlClient) Query(sqlStr string) []map[string]string {
	res, err := pgdb.Client.Query(sqlStr)
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

func (pgdb *PostgresqlClient) InsertOrUpdate(sqlStr string) int64 {
	result, err := pgdb.Client.Exec(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	return rows
}

func (pgdb *PostgresqlClient) Delete(sqlStr string) bool {
	res, err := pgdb.Client.Exec(sqlStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, _ := res.RowsAffected()
	log.Printf("delete ok!!!")
	if rows > 0 {
		return true
	} else {
		return false
	}
}
