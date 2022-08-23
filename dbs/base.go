package dbs

type SQLBase interface {
	Query()
	Close()
}
