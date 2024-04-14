package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"server/db/pg"
)

type Resolver struct{
	//any data source repository that implements the correct interface would work
	Repo pg.Repository
}
