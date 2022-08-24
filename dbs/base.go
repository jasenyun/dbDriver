package dbs

type SQLBase interface {
	Query(sqlStr string) []map[string]string
	InsertOrUpdate(sqlStr string) int64
	Delete(sqlStr string) bool
	Close()
}
