package connect

import (
	"context"
	//"fmt"
	"sync"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce sync.Once
)

func NewPgConnection(ctx context.Context, connection string) (*Postgres){
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connection)
		if err != nil {
			log.Fatal(err)
		}
		pgInstance = &Postgres{db}
	})
	return pgInstance
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.Db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}