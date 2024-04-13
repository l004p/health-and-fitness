package core

import (

)

type BillStatus string

const (
	PENDING BillStatus = "PENDING"
	COMPLETE BillStatus = "COMPLETE"
	CANCELLED BillStatus = "CANCELLED"
)

type BillRepository interface {

}