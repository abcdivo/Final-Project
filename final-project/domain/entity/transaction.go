package entity

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Transaction struct {
	transactionID    string
	customerID       string
	customerName     string
	revenue          int
	couponID         string
	discountPrice    int
	purchaseDate     time.Time
	transactionItems []*TransactionItems
}

type DTOTransaction struct {
	TransactionID    string
	CustomerID       string
	CustomerName     string
	Revenue          int
	CouponID         string
	DiscountPrice    int
	PurchaseDate     string
	TransactionItems []*TransactionItems
}

func (tr *Transaction) AddDataIDCustomer(customer *Customer) *Transaction {
	tr.customerID = customer.customerID

	return tr
}

func NewTransaction(dto DTOTransaction) (*Transaction, error) {
	if dto.CustomerID == "" {
		return nil, errors.New("customer ID cannot be empty")
	}
	// if dto.Revenue == 0 {
	// 	return nil, errors.New("revenue cannot be empty")
	// }
	if dto.PurchaseDate == "" {
		return nil, errors.New("purchase date cannot be empty")
	}

	convertPurchaseDate, _ := time.Parse("2006-01-02", dto.PurchaseDate)

	transaction := &Transaction{
		transactionID:    dto.TransactionID,
		customerID:       dto.CustomerID,
		customerName:     dto.CustomerName,
		revenue:          dto.Revenue,
		couponID:         dto.CouponID,
		discountPrice:    dto.DiscountPrice,
		purchaseDate:     convertPurchaseDate,
		transactionItems: dto.TransactionItems,
	}

	return transaction, nil
}

func (tr *Transaction) SetUniqTransactionID() *Transaction {
	rand.Seed(time.Now().UnixNano())
	min := 10000000
	max := 99999999
	valueString := strconv.Itoa(rand.Intn(max-min+1) + min)
	tr.transactionID = "TRAX" + valueString

	return tr
}

func (tr *Transaction) GetTransactionID() string {
	return tr.transactionID
}

func (tr *Transaction) GetCustomerID() string {
	return tr.customerID
}

func (tr *Transaction) GetRevenue() int {
	return tr.revenue
}

func (tr *Transaction) GetCouponID() string {
	return tr.couponID
}

func (tr *Transaction) GetPurchaseDate() string {
	return tr.purchaseDate.Format("2006-01-02")
}

func (tr *Transaction) GetTransactionItems() []*TransactionItems {
	return tr.transactionItems
}

func (tr *Transaction) SetTransactionItems(transactionItems []*TransactionItems) *Transaction {
	tr.transactionItems = transactionItems
	return tr
}

func (tr *Transaction) GetCustomerName() string {
	return tr.customerName
}

func (tr *Transaction) SumTotalRevenue(types string) (int, int) {
	var totalRevenue int
	var revenuePerCriteria int
	var dicountPrice int
	if types != "" {
		for _, item := range tr.transactionItems {
			if types == "ULTI" && (item.GetCriteriaID() == 3 || item.GetCriteriaID() == 5) {
				revenuePerCriteria += item.GetRevenueItem()
			} else if types == "PREMI" && item.GetCriteriaID() == 1 {
				revenuePerCriteria += item.GetRevenueItem()
			} else if types == "BASIC" && item.GetCriteriaID() == 4 {
				revenuePerCriteria += item.GetRevenueItem()
			} else {
				totalRevenue += item.GetRevenueItem()
			}
		}
	} else {
		for _, item := range tr.transactionItems {
			totalRevenue += item.GetRevenueItem()
		}
	}
	if revenuePerCriteria != 0 && types != "" {
		if types == "ULTI" {
			dicountPrice = int(float32(revenuePerCriteria) * 0.30)
		} else if types == "PREMI" {
			dicountPrice = int(float32(revenuePerCriteria) * 0.15)
		} else if types == "BASIC" {
			dicountPrice = int(float32(revenuePerCriteria) * 0.05)
		}
		revenuePerCriteria = revenuePerCriteria - dicountPrice
		totalRevenue = totalRevenue + revenuePerCriteria
	}
	tr.discountPrice = dicountPrice
	tr.revenue = totalRevenue
	return totalRevenue, dicountPrice
}

func (tr *Transaction) GetDiscountPrice() int {
	return tr.discountPrice
}
