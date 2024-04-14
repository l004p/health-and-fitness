package pg

import (
	"context"
	"server/core"
	"server/db/connect"
	//"github.com/jackc/pgx/v5"
	"fmt"
	"log"
	//"strconv"
)

type Repository interface {
	core.UserRepository
	core.MemberRepository
	core.TrainerRepository
	core.MembershipRepository
	core.EquipmentRepository
	core.BillRepository
	core.RoomRepository
}

type repoService struct {
	*Queries 
	db *connect.Postgres
}

func NewRepository(db *connect.Postgres) Repository {
	return &repoService{
		Queries: New(db.Db),
		db: db,
	}
}

func (r *repoService) TestFunction(ctx context.Context) int {
	users, err := r.getAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
			fmt.Printf("%v\n", user)
		}
	return 1
}