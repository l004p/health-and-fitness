package billservice

import(
	"context"
	"server/core"
)

func GetUserBills(r core.BillRepository, ctx context.Context, userID string) ([]core.Bill, error){
	b, err := r.GetUserBills(ctx, userID)
	if err != nil {
		return nil, err
	}
	return b, nil
}




func PayBill(r core.BillRepository, ctx context.Context, billID string) (core.Bill, error){
	//panic("not implemented")
	b, err := r.PayBill(ctx, billID)
	if err != nil {
		return core.Bill{}, err
	}
	return b, nil
}

func ChangeBillStatus(r core.BillRepository, ctx context.Context, billID string) (core.Bill, error){
	panic("not implemented")
}