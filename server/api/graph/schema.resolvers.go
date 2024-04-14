package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"server/api/graph/model"
	"server/api/middleware"
	"server/core"
	billservice "server/services/bill"
	equipmentservice "server/services/equipment"
	memberservice "server/services/member"
	roomservice "server/services/room"
	trainerservice "server/services/trainer"
	userservice "server/services/user"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// RegisterForEvent is the resolver for the RegisterForEvent field.
func (r *mutationResolver) RegisterForEvent(ctx context.Context, input model.Register) (bool, error) {
	panic(fmt.Errorf("not implemented: RegisterForEvent - RegisterForEvent"))
}

// UnregisterForEvent is the resolver for the UnregisterForEvent field.
func (r *mutationResolver) UnregisterForEvent(ctx context.Context, input model.Register) (bool, error) {
	panic(fmt.Errorf("not implemented: UnregisterForEvent - UnregisterForEvent"))
}

// ActivateEquipment is the resolver for the ActivateEquipment field.
func (r *mutationResolver) ActivateEquipment(ctx context.Context, input model.ActivateEquipment) (*model.Equipment, error) {
	//panic(fmt.Errorf("not implemented: ActivateEquipment - ActivateEquipment"))
	e, err := equipmentservice.ActivateStatus(r.Repo, ctx, input.EquipmentID, input.RoomID)
	//e, err := r.Repo.ActivateStatus(ctx, input.EquipmentID, input.RoomID)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	return &model.Equipment{
		EquipmentID:     e.EquipmentID,
		EquipmentType:   string(e.EquipmentType),
		EquipmentStatus: string(e.EquipmentStatus),
	}, nil
}

// DeactivateEquipment is the resolver for the DeactivateEquipment field.
func (r *mutationResolver) DeactivateEquipment(ctx context.Context, input model.DeactivateEquipment) (*model.Equipment, error) {
	//panic(fmt.Errorf("not implemented: DeactivateEquipment - DeactivateEquipment"))
	e, err := equipmentservice.DeactivateStatus(r.Repo, ctx, input.EquipmentID, core.EquipmentStatus(input.EquipmentStatus))

	//e, err := r.Repo.DeactivateStatus(ctx, input.EquipmentID, core.EquipmentStatus(input.EquipmentStatus))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	return &model.Equipment{
		EquipmentID:     e.EquipmentID,
		EquipmentType:   string(e.EquipmentType),
		EquipmentStatus: string(e.EquipmentStatus),
		Room:            nil,
	}, nil
}

// MoveEquipment is the resolver for the MoveEquipment field.
func (r *mutationResolver) MoveEquipment(ctx context.Context, input model.MoveEquipment) (bool, error) {
	//panic(fmt.Errorf("not implemented: MoveEquipment - MoveEquipment"))

	err := equipmentservice.AssignRoom(r.Repo, ctx, input.RoomID, input.EquipmentID)
	//err := r.Repo.MoveEquipment(ctx, input.RoomID, input.EquipmentID)
	if err != nil {
		return false, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	return true, nil
}

// CreateClass is the resolver for the CreateClass field.
func (r *mutationResolver) CreateClass(ctx context.Context, input model.Class) (bool, error) {
	panic(fmt.Errorf("not implemented: CreateClass - CreateClass"))
}

// EditClassStatus is the resolver for the EditClassStatus field.
func (r *mutationResolver) EditClassStatus(ctx context.Context, input model.ClassStatus) (bool, error) {
	panic(fmt.Errorf("not implemented: EditClassStatus - EditClassStatus"))
}

// AddClassTrainer is the resolver for the AddClassTrainer field.
func (r *mutationResolver) AddClassTrainer(ctx context.Context, input model.ClassTrainer) (bool, error) {
	panic(fmt.Errorf("not implemented: AddClassTrainer - AddClassTrainer"))
}

// RemoveClassTrainer is the resolver for the RemoveClassTrainer field.
func (r *mutationResolver) RemoveClassTrainer(ctx context.Context, input model.ClassTrainer) (bool, error) {
	panic(fmt.Errorf("not implemented: RemoveClassTrainer - RemoveClassTrainer"))
}

// CancelClass is the resolver for the CancelClass field.
func (r *mutationResolver) CancelClass(ctx context.Context, input model.ClassStatus) (bool, error) {
	panic(fmt.Errorf("not implemented: CancelClass - CancelClass"))
}

// CreateAvailableSession is the resolver for the CreateAvailableSession field.
func (r *mutationResolver) CreateAvailableSession(ctx context.Context, input model.Session) (bool, error) {
	panic(fmt.Errorf("not implemented: CreateAvailableSession - CreateAvailableSession"))
}

// EditSessionStatus is the resolver for the EditSessionStatus field.
func (r *mutationResolver) EditSessionStatus(ctx context.Context, input model.SessionStatus) (bool, error) {
	panic(fmt.Errorf("not implemented: EditSessionStatus - EditSessionStatus"))
}

// CancelSession is the resolver for the CancelSession field.
func (r *mutationResolver) CancelSession(ctx context.Context, input model.SessionStatus) (bool, error) {
	panic(fmt.Errorf("not implemented: CancelSession - CancelSession"))
}

// PayBill is the resolver for the PayBill field.
func (r *mutationResolver) PayBill(ctx context.Context, input model.PayBill) (bool, error) {
	//panic(fmt.Errorf("not implemented: PayBill - PayBill"))
	_, err := billservice.PayBill(r.Repo, ctx, input.BillID)
	if err != nil {
		return false, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	return true, nil
}

// UpdateBillStatus is the resolver for the UpdateBillStatus field.
func (r *mutationResolver) UpdateBillStatus(ctx context.Context, input model.PayBill) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateBillStatus - UpdateBillStatus"))
}

// AddMetric is the resolver for the AddMetric field.
func (r *mutationResolver) AddMetric(ctx context.Context, input model.NewMetric) (bool, error) {
	panic(fmt.Errorf("not implemented: AddMetric - AddMetric"))
}

// AddGoal is the resolver for the AddGoal field.
func (r *mutationResolver) AddGoal(ctx context.Context, input model.NewGoal) (bool, error) {
	panic(fmt.Errorf("not implemented: AddGoal - AddGoal"))
}

// ViewProfile is the resolver for the ViewProfile field.
func (r *queryResolver) ViewProfile(ctx context.Context) (*model.Profile, error) {
	//panic(fmt.Errorf("not implemented: ViewProfile - ViewProfile"))
	getID := ctx.Value(middleware.CtxKey("userID")).(string)
	//fmt.Printf("HERE %s", getID)
	thisUser, err := userservice.UserByID(r.Repo, ctx, getID)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}

	user := &model.User{
		UserID:    getID,
		Username:  thisUser.Username,
		UserEmail: thisUser.UserEmail,
		FirstName: thisUser.FirstName,
		LastName:  thisUser.LastName,
	}

	//TODO: implement this stuff
	membership := &model.Membership{}

	var goals []*model.Goal

	var metrics []*model.Metric

	profile := &model.Profile{
		User:       user,
		Membership: membership,
		Goals:      goals,
		Metrics:    metrics,
	}

	return profile, nil
}

// ViewUpcomingEvents is the resolver for the ViewUpcomingEvents field.
func (r *queryResolver) ViewUpcomingEvents(ctx context.Context, typeArg *string) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented: ViewUpcomingEvents - ViewUpcomingEvents"))
}

// ViewUpcomingEventsLead is the resolver for the ViewUpcomingEventsLead field.
func (r *queryResolver) ViewUpcomingEventsLead(ctx context.Context, typeArg *string) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented: ViewUpcomingEventsLead - ViewUpcomingEventsLead"))
}

// GetAllUsersTrained is the resolver for the GetAllUsersTrained field.
func (r *queryResolver) GetAllUsersTrained(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: GetAllUsersTrained - GetAllUsersTrained"))
}

// SearchMember is the resolver for the SearchMember field.
func (r *queryResolver) SearchMember(ctx context.Context, firstname *string, lastname *string) ([]*model.Profile, error) {
	//panic(fmt.Errorf("not implemented: SearchMember - SearchMember"))
	var fn string
	var ln string
	if firstname != nil {
		fn = *firstname
	}
	if lastname != nil {
		ln = *lastname
	}
	u, err := trainerservice.SearchForUser(r.Repo, ctx, fn, ln)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	var profiles []*model.Profile
	for _, current := range u {
		//fmt.Println(current.Username)
		model := &model.Profile{
			User: &model.User{
				Username: current.Username,
				FirstName: current.FirstName,
				LastName: current.LastName,
			},
		}

		profiles = append(profiles, model)
	}
	return profiles, nil
}

// GetTrainers is the resolver for the GetTrainers field.
func (r *queryResolver) GetTrainers(ctx context.Context) ([]*model.User, error) {
	//panic(fmt.Errorf("not implemented: GetTrainers - GetTrainers"))
	getID := ctx.Value(middleware.CtxKey("userID")).(string)
	t, err := memberservice.GetTrainers(r.Repo, ctx, getID)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	var result []*model.User
	for _, current := range t {
		model := &model.User{
			Username:  current.Username,
			FirstName: current.FirstName,
			LastName:  current.LastName,
		}
		result = append(result, model)
	}
	return result, nil
}

// GetBills is the resolver for the GetBills field.
func (r *queryResolver) GetBills(ctx context.Context, status *string) ([]*model.Bill, error) {
	//panic(fmt.Errorf("not implemented: GetBills - GetBills"))
	var bills []*model.Bill
	getID := ctx.Value(middleware.CtxKey("userID")).(string)
	if status == nil {
		b, err := billservice.GetUserBills(r.Repo, ctx, getID)
		if err != nil {
			return nil, &gqlerror.Error{
				Message: err.Error(),
			}
		}
		for _, current := range b {
			model := &model.Bill{
				BillID:      current.BillID,
				UserID:      getID,
				Description: current.Description,
				Status:      string(current.Status),
				Date:        current.Date.String(),
			}
			bills = append(bills, model)
		}
	}
	return bills, nil
}

// ViewAllBills is the resolver for the ViewAllBills field.
func (r *queryResolver) ViewAllBills(ctx context.Context, status *string, id *string) ([]*model.Bill, error) {
	var bills []*model.Bill
	//getID := ctx.Value(middleware.CtxKey("userID")).(string)
	if status == nil {
		b, err := billservice.GetUserBills(r.Repo, ctx, *id)
		if err != nil {
			return nil, &gqlerror.Error{
				Message: err.Error(),
			}
		}
		for _, current := range b {
			model := &model.Bill{
				BillID:      current.BillID,
				UserID:      current.UserID,
				Description: current.Description,
				Status:      string(current.Status),
				Date:        current.Date.String(),
			}
			bills = append(bills, model)
		}
	}
	return bills, nil
}

// GetAllEquipment is the resolver for the GetAllEquipment field.
func (r *queryResolver) GetAllEquipment(ctx context.Context, status *string, typeArg *string) ([]*model.Equipment, error) {
	//panic(fmt.Errorf("not implemented: GetAllEquipment - GetAllEquipment"))
	var equipment []*model.Equipment

	if status == nil && typeArg == nil {
		e, err := equipmentservice.Equipment(r.Repo, ctx)
		//e, err := r.Repo.AllEquipment(ctx)
		if err != nil {
			return nil, &gqlerror.Error{
				Message: err.Error(),
			}
		}

		for _, current := range e {
			model := &model.Equipment{
				EquipmentID:     current.EquipmentID,
				EquipmentType:   string(current.EquipmentType),
				EquipmentStatus: string(current.EquipmentStatus),
				Room: &model.Room{
					RoomID:      current.Room.RoomID,
					Description: current.Room.Description,
				},
			}
			equipment = append(equipment, model)
		}
	} else if status != nil && typeArg == nil {
		e, err := equipmentservice.ByStatus(r.Repo, ctx, core.EquipmentStatus(*status))
		//e, err := r.Repo.EquipmentByStatus(ctx, core.EquipmentStatus(*status))
		if err != nil {
			return nil, &gqlerror.Error{
				Message: err.Error(),
			}
		}

		for _, current := range e {
			model := &model.Equipment{
				EquipmentID:     current.EquipmentID,
				EquipmentType:   string(current.EquipmentType),
				EquipmentStatus: string(current.EquipmentStatus),
				Room: &model.Room{
					RoomID:      current.Room.RoomID,
					Description: current.Room.Description,
				},
			}
			equipment = append(equipment, model)
		}
	} else if status == nil && typeArg != nil {
		e, err := equipmentservice.ByType(r.Repo, ctx, core.EquipmentType(*typeArg))
		//e, err := r.Repo.EquipmentByType(ctx, core.EquipmentType(*typeArg))
		if err != nil {
			return nil, &gqlerror.Error{
				Message: err.Error(),
			}
		}

		for _, current := range e {
			model := &model.Equipment{
				EquipmentID:     current.EquipmentID,
				EquipmentType:   string(current.EquipmentType),
				EquipmentStatus: string(current.EquipmentStatus),
				Room: &model.Room{
					RoomID:      current.Room.RoomID,
					Description: current.Room.Description,
				},
			}
			equipment = append(equipment, model)
		}
	} else {
		e, err := equipmentservice.ByStatusAndType(r.Repo, ctx, core.EquipmentStatus(*status), core.EquipmentType(*typeArg))
		//e, err := r.Repo.EquipmentByStatusAndType(ctx, core.EquipmentType(*typeArg), core.EquipmentStatus(*status))
		if err != nil {
			return nil, &gqlerror.Error{
				Message: err.Error(),
			}
		}

		for _, current := range e {
			model := &model.Equipment{
				EquipmentID:     current.EquipmentID,
				EquipmentType:   string(current.EquipmentType),
				EquipmentStatus: string(current.EquipmentStatus),
				Room: &model.Room{
					RoomID:      current.Room.RoomID,
					Description: current.Room.Description,
				},
			}
			equipment = append(equipment, model)
		}
	}

	return equipment, nil
}

// GetEquipment is the resolver for the GetEquipment field.
func (r *queryResolver) GetEquipment(ctx context.Context, id string) (*model.Equipment, error) {
	//panic(fmt.Errorf("not implemented: GetEquipment - GetEquipment"))
	e, err := equipmentservice.EquipmentByID(r.Repo, ctx, id)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	//TODO: room

	return &model.Equipment{
		EquipmentID:     e.EquipmentID,
		EquipmentType:   string(e.EquipmentType),
		EquipmentStatus: string(e.EquipmentStatus),
		Room:            nil,
	}, nil
}

// GetRooms is the resolver for the GetRooms field.
func (r *queryResolver) GetRooms(ctx context.Context) ([]*model.Room, error) {
	//panic(fmt.Errorf("not implemented: GetRooms - GetRooms"))
	rooms, err := roomservice.GetAllRooms(r.Repo, ctx)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	var result []*model.Room
	for _, current := range rooms {
		model := &model.Room{
			RoomID:      current.RoomID,
			Description: current.Description,
		}
		result = append(result, model)
	}
	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//panic(fmt.Errorf("not implemented: Users - users"))
	user1 := &model.User{
		UserID:    strconv.FormatInt(int64(1), 10),
		Username:  "testName",
		FirstName: "Justin",
		LastName:  "Foley",
	}
	user2 := &model.User{
		UserID:    strconv.FormatInt(int64(2), 10),
		Username:  "testName2",
		FirstName: "Jessica",
		LastName:  "Davies",
	}
	var users []*model.User
	users = append(users, user1, user2)
	userservice.TestFunction(r.Repo, ctx)
	return users, nil
}
