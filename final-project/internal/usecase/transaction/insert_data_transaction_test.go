package transaction_test

import (
	"context"
	"final-project/domain/entity"
	databasesql "final-project/internal/config/database/mysql"
	"final-project/internal/repository/mysql"
	"final-project/internal/usecase/transaction"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsecaseInsertDataTransaction(t *testing.T) {
	var (
		ctx                             = context.Background()
		connectionDatabase              = databasesql.InitMysqlDB()
		transactionRepositoryMysql      = mysql.NewTransactionMysql(connectionDatabase)
		transactionItemsRepositoryMysql = mysql.NewTransactionItemsMysql(connectionDatabase)
		couponRepositoryMysql           = mysql.NewCouponMysql(connectionDatabase)
	)

	transactionUsecase := transaction.NewTransactionUsecase(transactionRepositoryMysql, transactionItemsRepositoryMysql, couponRepositoryMysql)

	listItems := make([]*entity.TransactionItems, 0)

	item1, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		CriteriaID:  2,
		RevenueItem: 500000,
	})

	item2, _ := entity.NewTransactionItems(entity.DTOTransactionItems{
		CriteriaID:  3,
		RevenueItem: 1000000,
	})

	listItems = append(listItems, item1, item2)

	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerID:       "CUST10293847",
		Revenue:          1500000,
		CouponID:         "",
		PurchaseDate:     "2022-12-12",
		TransactionItems: listItems,
	})

	transaction.SetUniqTransactionID()

	if err != nil {
		fmt.Println(err)
	}

	_, errInsert := transactionUsecase.InsertDataTransaction(ctx, transaction)

	assert.Nil(t, errInsert)
}
