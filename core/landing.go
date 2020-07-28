package core

import "context"

type Landing struct {
	UniqueUsersPerMonthActual int
	UniqueUsersPerMonthShift  int

	AvgSiteReachActual int
	AvgSiteReachShift  int

	AvgChannelReachActual int
	AvgChannelReachShift  int
}

func (l Landing) UniqueUsersPerMonth() int {
	return l.UniqueUsersPerMonthShift + l.UniqueUsersPerMonthActual
}

func (l Landing) AvgSiteReach() int {
	return l.AvgSiteReachShift + l.AvgSiteReachActual
}

func (l Landing) AvgChannelReach() int {
	return l.AvgChannelReachShift + l.AvgChannelReachActual
}

type LandingStore interface {
	Get(ctx context.Context) (*Landing, error)
	Update(ctx context.Context, landing *Landing) error
}
