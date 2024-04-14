package core

import (
	"context"
	"time"
)

type BillStatus string

const (
	PENDING BillStatus = "PENDING"
	COMPLETE BillStatus = "COMPLETE"
	CANCELLED BillStatus = "CANCELLED"
)

type Bill struct {
	BillID string
	UserID string
	Description string
	Status BillStatus
	Date time.Time

}

type BillRepository interface {
	GetUserBills(ctx context.Context, userID string) ([]Bill, error)
	PayBill(ctx context.Context, billID string) (Bill, error)
	ChangeBillStatus(ctx context.Context, billID string) (Bill, error)
}