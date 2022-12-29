package mysql_test

import (
	"context"
	databasesql "final-project/internal/config/database/mysql"
	"final-project/internal/repository/mysql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepoSumTransactionById(t *testing.T) {
	var (
		ctx                = context.Background()
		connectionDatabase = databasesql.InitMysqlDB()
		repoTransaction    = mysql.NewTransactionMysql(connectionDatabase)
		customer_id        = "CUST15399180"
		// lastDateCreated    = "2022-12-23 13:30:29"
		lastDateCreated = ""
	)

	sumRevenue := repoTransaction.SumTransactionById(ctx, customer_id, lastDateCreated)

	fmt.Println(sumRevenue)
	assert.NotEqual(t, sumRevenue, 0)
}
