package mysql

import (
	"context"
	"database/sql"
	"errors"
	"final-project/domain/entity"
	"time"
)

type TransactionMysqlInteractor struct {
	db *sql.DB
}

func NewTransactionMysql(db *sql.DB) *TransactionMysqlInteractor {
	return &TransactionMysqlInteractor{
		db: db,
	}
}

func (m *TransactionMysqlInteractor) InsertDataTransaction(ctx context.Context, dataTransaction *entity.Transaction) (string, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO transaction (transaction_id, customer_id, revenue, coupon_id, discount_price, purchase_date)" +
		"VALUES(?, ?, ?, ?, ?, ?)"

	_, errMysql = m.db.Exec(insertQuery, dataTransaction.GetTransactionID(), dataTransaction.GetCustomerID(),
		dataTransaction.GetRevenue(), dataTransaction.GetCouponID(), dataTransaction.GetDiscountPrice(), dataTransaction.GetPurchaseDate())

	if errMysql != nil {
		return "", errMysql
	}

	return dataTransaction.GetTransactionID(), nil
}

func (m *TransactionMysqlInteractor) GetListTransaction(ctx context.Context, limit int) ([]*entity.Transaction, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT transaction_id, transaction.customer_id, customer.name, revenue, coupon_id, purchase_date FROM transaction " +
		"LEFT JOIN customer ON transaction.customer_id = customer.customer_id ORDER BY purchase_date LIMIT ?"
	rows, errMysql := m.db.QueryContext(ctx, sqlQuery, limit)
	if errMysql != nil {
		return nil, errMysql
	}

	transactionCollection := make([]*entity.Transaction, 0)
	for rows.Next() {
		var (
			transaction_id string
			customer_id    string
			customer_name  string
			revenue        int
			coupon_id      string
			purchase_date  string
		)

		errTransaction := rows.Scan(&transaction_id, &customer_id, &customer_name, &revenue, &coupon_id, &purchase_date)

		if errTransaction != nil {
			return nil, errTransaction
		}

		dateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", purchase_date)
		if errParse != nil {
			return nil, errParse
		}
		purchase_date = dateParse.Format("2006-01-02")

		transaction, errNewTransaction := entity.NewTransaction(entity.DTOTransaction{
			TransactionID: transaction_id,
			CustomerID:    customer_id,
			CustomerName:  customer_name,
			Revenue:       revenue,
			CouponID:      coupon_id,
			PurchaseDate:  purchase_date,
		})

		if errNewTransaction != nil {
			return nil, errNewTransaction
		}
		transactionCollection = append(transactionCollection, transaction)
	}
	defer rows.Close()

	return transactionCollection, nil
}

func (m *TransactionMysqlInteractor) DeleteTransactionById(ctx context.Context, transaction_id string) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM transaction where transaction_id = ?"
	row, errMysql := m.db.ExecContext(ctx, deleteQuery, transaction_id)

	if errMysql != nil && errMysql != sql.ErrNoRows {
		return errMysql
	}

	check, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if check != 1 {
		return errors.New("transaction id not found")
	}

	return nil
}

func (m *TransactionMysqlInteractor) SumTransactionById(ctx context.Context, customer_id string, lastDateCreated string) int64 {
	var (
		errMysql    error
		sum_revenue int64
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	if lastDateCreated != "" {
		sqlQuery := "SELECT SUM(revenue) as sum_revenue FROM transaction WHERE customer_id = ? AND purchase_date > ?"
		errMysql = m.db.QueryRowContext(ctx, sqlQuery, customer_id, lastDateCreated).Scan(&sum_revenue)
	} else {
		sqlQuery := "SELECT SUM(revenue) as sum_revenue FROM transaction WHERE customer_id = ?"
		errMysql = m.db.QueryRowContext(ctx, sqlQuery, customer_id).Scan(&sum_revenue)
	}

	if errMysql == sql.ErrNoRows {
		return 0
	}

	if errMysql != nil {
		return 0
	}

	return sum_revenue
}
