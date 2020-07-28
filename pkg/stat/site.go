package stat

import "context"

type Site interface {
	GetUniqueVisitorsPerMonth(ctx context.Context) (int, error)
}
