package mysql

import (
	"context"
	"database/sql"
	"errors"
	"final-project/domain/entity"
	"time"
)

type TransactionItemsMysqlInteractor struct {
	db *sql.DB
}

func NewTransactionItemsMysql(db *sql.DB) *TransactionItemsMysqlInteractor {
	return &TransactionItemsMysqlInteractor{
		db: db,
	}
}

func (m *TransactionItemsMysqlInteractor) InsertDataTransactionItems(ctx context.Context, dataTransactionItems *entity.TransactionItems, transactionID string) error {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO transaction_items (transaction_id, criteria_id, revenue_item, date_created)" +
		"VALUES(?, ?, ?, ?)"

	_, errMysql = m.db.Exec(insertQuery, transactionID, dataTransactionItems.GetCriteriaID(),
		dataTransactionItems.GetRevenueItem(), dataTransactionItems.GetDateCreated())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

func (m *TransactionItemsMysqlInteractor) GetItemsByTransactionID(ctx context.Context, transaction_id string) ([]*entity.TransactionItems, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	sqlQuery := "SELECT item_id, transaction_id, transaction_items.criteria_id, criteria.criteria_name, revenue_item, date_created FROM transaction_items " +
		" LEFT JOIN criteria ON transaction_items.criteria_id = criteria.criteria_id WHERE transaction_id = ?"
	rows, errMysql := m.db.QueryContext(ctx, sqlQuery, transaction_id)
	if errMysql != nil {
		return nil, errMysql
	}

	itemsCollection := make([]*entity.TransactionItems, 0)
	for rows.Next() {
		var (
			item_id       int
			transactionID string
			criteria_id   int
			criteria_name string
			revenue_item  int
			date_created  string
		)

		errTransaction := rows.Scan(&item_id, &transactionID, &criteria_id, &criteria_name, &revenue_item, &date_created)

		if errTransaction != nil {
			return nil, errTransaction
		}
		item, errNewItem := entity.NewTransactionItems(entity.DTOTransactionItems{
			ItemID:        item_id,
			TransactionID: transactionID,
			CriteriaID:    criteria_id,
			CriteriaName:  criteria_name,
			RevenueItem:   revenue_item,
			DateCreated:   date_created,
		})

		if errNewItem != nil {
			return nil, errNewItem
		}
		itemsCollection = append(itemsCollection, item)
	}
	defer rows.Close()

	return itemsCollection, nil
}

func (m *TransactionItemsMysqlInteractor) DeleteItemsByTransactionId(ctx context.Context, transaction_id string) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	deleteQuery := "DELETE FROM transaction_items where transaction_id = ?"
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
