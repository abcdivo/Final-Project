package mysql

import (
	"context"
	"database/sql"
	"errors"
	"final-project/domain/entity"
	"time"
)

type CouponMysqlInteractor struct {
	db *sql.DB
}

func NewCouponMysql(db *sql.DB) *CouponMysqlInteractor {
	return &CouponMysqlInteractor{
		db: db,
	}
}

func (c *CouponMysqlInteractor) GetCouponById(ctx context.Context, id_coupon string) (*entity.Coupon, error) {
	var (
		errMysql     error
		coupon_id    string
		types        string
		expired_date string
		customer_id  string
		status       int
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT coupon_id, types, expired_date, customer_id, status FROM coupon WHERE coupon_id = ?"
	errMysql = c.db.QueryRowContext(ctx, sqlQuery, id_coupon).Scan(&coupon_id, &types, &expired_date, &customer_id, &status)

	if errMysql != nil {
		return nil, errMysql
	}

	dateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", expired_date)
	if errParse != nil {
		return nil, errParse
	}
	expired_date = dateParse.Format("2006-01-02")

	coupon, errCoupon := entity.NewCoupon(entity.DTOCoupon{
		CouponID:    coupon_id,
		Types:       types,
		CustomerID:  customer_id,
		ExpiredDate: expired_date,
		Status:      status,
	})

	if errCoupon != nil {
		return nil, errCoupon
	}

	return coupon, nil
}

func (c *CouponMysqlInteractor) InsertDataCoupon(ctx context.Context, dataCoupon *entity.Coupon) (string, error) {
	var (
		errMysql error
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO coupon (coupon_id, types, expired_date, customer_id, status, created_date)" +
		"VALUES(?, ?, ?, ?, ?, ?)"

	_, errMysql = c.db.Exec(insertQuery, dataCoupon.GetCouponID(), dataCoupon.GetTypes(), dataCoupon.GetExpiredDate(), dataCoupon.GetCustomerID(), dataCoupon.GetStatus(), dataCoupon.GetDateCreated())

	if errMysql != nil {
		return "", errMysql
	}
	return dataCoupon.GetCouponID(), nil
}

func (c *CouponMysqlInteractor) CouponValidation(ctx context.Context, dataTransaction *entity.Transaction) (string, error) {
	if dataTransaction.GetCouponID() != "" {
		coupon, errCoupon := c.GetCouponById(ctx, dataTransaction.GetCouponID())
		if errCoupon != nil {
			if errCoupon == sql.ErrNoRows {
				return "", errors.New("coupon not found")
			}
			return "", errCoupon
		}

		if coupon.GetCustomerID() != dataTransaction.GetCustomerID() {
			return "", errors.New("this coupon not eligible for this customer")
		}

		if coupon.GetStatus() == 1 {
			return "", errors.New("this coupon was used")
		}

		today := time.Now().Format("2006-01-02")
		if coupon.GetExpiredDate() < today {
			return "", errors.New("coupon expired")
		}

		return coupon.GetTypes(), nil
	}

	return "", nil
}

func (c *CouponMysqlInteractor) GetCouponByCustomerId(ctx context.Context, customer_id string) ([]*entity.Coupon, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT coupon_id, types, expired_date, status, customer_id FROM coupon WHERE customer_id = ?"
	rows, errMysql := c.db.QueryContext(ctx, sqlQuery, customer_id)
	if errMysql != nil {
		return nil, errMysql
	}

	couponCollection := make([]*entity.Coupon, 0)
	for rows.Next() {
		var (
			coupon_id    string
			types        string
			expired_date string
			status       int
			customer_id  string
		)

		errCoupon := rows.Scan(&coupon_id, &types, &expired_date, &status, &customer_id)

		if errCoupon != nil {
			return nil, errCoupon
		}
		dateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", expired_date)
		if errParse != nil {
			return nil, errParse
		}
		expired_date = dateParse.Format("2006-01-02")

		coupon, errNewCoupon := entity.NewCoupon(entity.DTOCoupon{
			CouponID:    coupon_id,
			Types:       types,
			ExpiredDate: expired_date,
			CustomerID:  customer_id,
			Status:      status,
		})
		if errNewCoupon != nil {
			return nil, errNewCoupon
		}
		couponCollection = append(couponCollection, coupon)
	}
	defer rows.Close()

	return couponCollection, nil
}

func (c *CouponMysqlInteractor) UpdateCouponStatus(ctx context.Context, coupon_id string, status int) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	result, err := c.db.ExecContext(ctx, "UPDATE coupon SET status = ? WHERE coupon_id = ?", status, coupon_id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("coupon status not uupdated")
	}

	return nil
}

func (c *CouponMysqlInteractor) GetLastCreatedCouponByCustomerId(ctx context.Context, customer_id string) (*entity.Coupon, error) {
	var (
		errMysql     error
		coupon_id    string
		types        string
		expired_date string
		customerID   string
		status       int
		created_date string
	)

	_, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT coupon_id, types, expired_date, customer_id, status, created_date FROM coupon WHERE customer_id = ? ORDER BY created_date DESC LIMIT 1"
	errMysql = c.db.QueryRowContext(ctx, sqlQuery, customer_id).Scan(&coupon_id, &types, &expired_date, &customerID, &status, &created_date)

	if errMysql == sql.ErrNoRows {
		return nil, nil
	}

	if errMysql != nil {
		return nil, errMysql
	}

	expDateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", expired_date)
	if errParse != nil {
		return nil, errParse
	}
	expired_date = expDateParse.Format("2006-01-02")

	createdDateParse, errParse := time.Parse("2006-01-02T15:04:05-07:00", created_date)
	if errParse != nil {
		return nil, errParse
	}
	created_date = createdDateParse.Format("2006-01-02 15:04:05")

	coupon, errCoupon := entity.NewCoupon(entity.DTOCoupon{
		CouponID:    coupon_id,
		Types:       types,
		CustomerID:  customer_id,
		ExpiredDate: expired_date,
		Status:      status,
		CreatedDate: created_date,
	})

	if errCoupon != nil {
		return nil, errCoupon
	}

	return coupon, nil
}
