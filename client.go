package dbDriver

import (
	"github.com/jasenyun/dbDriver/dbs"
	"log"
)

type DB_TYPE string

const (
	MYSQL_DB   DB_TYPE = "MYSQL_DB"
	ORACLE_DB          = "ORACLE_DB"
	POSTGRESQL         = "POSTGRESQL"
)

type DbClient struct {
	DbType       DB_TYPE
	MysqlDb      *dbs.MysqlClient
	OracleDb     *dbs.OracleClient
	PostgresqlDb *dbs.PostgresqlClient
}

var DbCli *DbClient

func InitInstance(dbType DB_TYPE, dns string) {
	DbCli = &DbClient{}
	DbCli.DbType = dbType
	switch dbType {
	case MYSQL_DB:
		DbCli.MysqlDb.Client = dbs.NewMySql(dns)
		break
	case ORACLE_DB:
		DbCli.MysqlDb.Client = dbs.NewOracle(dns)
		break
	case POSTGRESQL:
		DbCli.MysqlDb.Client = dbs.NewPostgresql(dns)
		break
	default:
		log.Fatal("未找到对应的数据库")
		break
	}
}

func (db *DbClient) getDriver() dbs.SQLBase {
	if db.MysqlDb.Client != nil {
		return db.MysqlDb
	}
	if db.OracleDb.Client != nil {
		return db.OracleDb
	}
	return db.PostgresqlDb
}

func (db *DbClient) Close() {
	db.getDriver().Close()
}
func (db *DbClient) Query(sqlStr string) []map[string]string {
	return db.getDriver().Query(sqlStr)
}

func (db *DbClient) InsertOrUpdate(sqlStr string) int64 {
	return db.InsertOrUpdate(sqlStr)
}

func (db *DbClient) Delete(sqlStr string) bool {
	return db.getDriver().Delete(sqlStr)
}
