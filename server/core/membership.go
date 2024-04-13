package core

import (
	"time"
	"context"
)


type MembershipType string

const (
	REGULAR MembershipType= "REGULAR"
	EXTENDED MembershipType = "EXTENDED"
)

type Status string

const (
	ACTIVE Status= "ACTIVE"
	INACTIVE Status = "INACTIVE"
)

type Membership struct {
	MembershipID int32
	UserID		 int32
	Status		 Status
	StartDate	 time.Time
	EndDate		 time.Time
}

type MembershipRepository interface {
	AddMembership(ctx context.Context, user User, membership MembershipType) (error)
	CancelMembership(ctx context.Context, user User, membership MembershipType) (error)
	ValidMembership(ctx context.Context, user User) (bool, error)
}