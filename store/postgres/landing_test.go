package postgres

import (
	"context"
	"testing"

	"github.com/bots-house/birzzha/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLandingStore(t *testing.T) {
	ctx, store := context.Background(), newPostgres(t).Landing

	t.Run("Get", func(t *testing.T) {
		landing, err := store.Get(ctx)
		require.NoError(t, err)
		require.NotNil(t, landing)

		assert.Equal(t, defaultLanding, landing)
	})

	t.Run("Update", func(t *testing.T) {
		excepted := &core.Landing{
			UniqueUsersPerMonthActual: 1,
			UniqueUsersPerMonthShift:  2,

			AvgSiteReachActual: 3,
			AvgSiteReachShift:  4,

			AvgChannelReachActual: 5,
			AvgChannelReachShift:  7,
		}

		err := store.Update(ctx, excepted)
		require.NoError(t, err)

		actual, err := store.Get(ctx)
		require.NoError(t, err)

		assert.Equal(t, excepted, actual)
	})

}
