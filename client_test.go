package dbDriver

import (
	"fmt"
	"testing"
)

func TestInitInstance(t *testing.T) {
	mysql_dns := "root:123456@tcp(127.0.0.1:3306)/testSQL"

	InitInstance(MYSQL_DB, mysql_dns)
	result := DbCli.Query("select * from users")
	fmt.Println(result)
}
