package dbDriver

import "github.com/jasenyun/dbDriver/dbs"

type DB_TYPE string

const (
	MYSQL_DB   DB_TYPE = "mysql"
	ORACLE_DB          = "oracle"
	POSTGRESQL         = "postgresql"
)

type DbClient struct {
	DbType       DB_TYPE
	MysqlDb      *dbs.MysqlClient
	OracleDb     *dbs.OracleClient
	PostgresqlDb *dbs.PostgresqlClient
}

var DbCli *DbClient

func InitInstance(dbType DB_TYPE) {
	switch dbType {
	case MYSQL_DB:
		DbCli.MysqlDb.Client = dbs.NewMySql()
		break
	case ORACLE_DB:
		DbCli.MysqlDb.Client = dbs.NewOracle()
		break
	case POSTGRESQL:
		DbCli.MysqlDb.Client = dbs.NewPostgresql()
		break
	default:
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
func (db *DbClient) Query() {
	db.getDriver().Close()
}
