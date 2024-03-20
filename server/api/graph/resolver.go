package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"server/db/pg"
)

type Resolver struct{
	Repo pg.Repository
}
