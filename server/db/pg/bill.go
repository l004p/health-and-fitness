package pg

import (
	"server/core"
	"context"
	"strconv"
)

func (r *repoService) GetUserBills(ctx context.Context, userID string) ([]core.Bill, error){
	//panic("not implemented")
	uID, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		return nil, err
	}
	b, err := r.getAllUserBills(ctx, int32(uID))
	if err != nil {
		return nil, err
	}
	var bills []core.Bill
	for _, bill := range b {
		//fmt.Println(piece.EquipmentID)
		current := core.Bill{
			BillID: strconv.Itoa(int(bill.BillID)),
			UserID: userID,
			Description: bill.BillDescription,
			Status: core.BillStatus(bill.BillStatus),
			Date: bill.BillDate.Time,
		}
		//fmt.Println("here")
		bills = append(bills, current)
	}
	return bills, nil
}


func (r *repoService) PayBill(ctx context.Context, billID string) (core.Bill, error){
	//panic("not implemented")
	idInt, err := strconv.ParseInt(billID, 10, 32)
	if err != nil {
		return core.Bill{}, err
	}
	err = r.payBill(ctx, int32(idInt))
	if err != nil {
		return core.Bill{}, err
	}
	bill, err := r.getBillByID(ctx, int32(idInt))
	if err != nil {
		return core.Bill{}, err
	}
	return core.Bill{
		BillID: strconv.Itoa(int(bill.BillID)),
		UserID: strconv.Itoa(int(bill.UserID)),
		Description: bill.BillDescription,
		Status: core.BillStatus(bill.BillStatus),
		Date: bill.BillDate.Time,
	}, nil
}

func (r *repoService) ChangeBillStatus(ctx context.Context, billID string) (core.Bill, error){
	panic("not implemented")
}