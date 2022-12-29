package mysql_test

import (
	"context"
	"errors"
	"final-project/domain/entity"
	databasesql "final-project/internal/config/database/mysql"
	"final-project/internal/repository/mysql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestScenarioCouponValidation struct {
	CustomerID   string
	Revenue      int
	CouponID     string
	PurchaseDate string
	Want         error
}

func TestRepoGetCouponById(t *testing.T) {
	var (
		ctx                   = context.Background()
		connectionDatabase    = databasesql.InitMysqlDB()
		couponRepositoryMysql = mysql.NewCouponMysql(connectionDatabase)
		id_coupon             = "PREMI-RND1209318092312"
	)

	coupon, err := couponRepositoryMysql.GetCouponById(ctx, id_coupon)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(coupon)
	assert.Nil(t, err)
}

func TestRepoCouponValidationNegativeTest(t *testing.T) {
	var (
		ctx                   = context.Background()
		connectionDatabase    = databasesql.InitMysqlDB()
		couponRepositoryMysql = mysql.NewCouponMysql(connectionDatabase)
	)

	listScenario := []TestScenarioCouponValidation{
		{
			CustomerID:   "CUST15399180",
			Revenue:      1500000,
			CouponID:     "PREMI-RND1209318092312",
			PurchaseDate: "2022-12-12",
			Want:         errors.New("coupon expired"),
		},
		{
			CustomerID:   "CUST15399180",
			Revenue:      1500000,
			CouponID:     "PREMI-RND1209318092315",
			PurchaseDate: "2022-12-12",
			Want:         errors.New("coupon not found"),
		},
		{
			CustomerID:   "CUST15399180",
			Revenue:      1500000,
			CouponID:     "ULTI-RND7821387123456",
			PurchaseDate: "2022-12-12",
			Want:         errors.New("this coupon not eligible for this customer"),
		},
		{
			CustomerID:   "CUST22456861",
			Revenue:      3000000,
			CouponID:     "BASIC-RND1923808132345",
			PurchaseDate: "2022-12-12",
			Want:         errors.New("this coupon was used"),
		},
	}

	for _, testCase := range listScenario {
		dataTransaction, _ := entity.NewTransaction(entity.DTOTransaction{
			CustomerID:   testCase.CustomerID,
			Revenue:      testCase.Revenue,
			CouponID:     testCase.CouponID,
			PurchaseDate: testCase.PurchaseDate,
		})
		dataTransaction.SetUniqTransactionID()
		_, errValidation := couponRepositoryMysql.CouponValidation(ctx, dataTransaction)
		assert.Equal(t, testCase.Want, errValidation)
	}
}

func TestRepoGetLastCreatedCouponByCustomerId(t *testing.T) {
	var (
		ctx                   = context.Background()
		connectionDatabase    = databasesql.InitMysqlDB()
		couponRepositoryMysql = mysql.NewCouponMysql(connectionDatabase)
		customer_id           = "CUST15399180"
	)

	coupon, err := couponRepositoryMysql.GetLastCreatedCouponByCustomerId(ctx, customer_id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(coupon.GetDateCreated())
	fmt.Printf("date_created: %T\n", coupon.GetDateCreated())
	assert.Nil(t, err)
}
