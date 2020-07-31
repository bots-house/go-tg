package catalog

import (
	"context"
	"fmt"
	"testing"

	"github.com/bots-house/birzzha/pkg/stat"
	"github.com/stretchr/testify/assert"
)

func TestService_GetDailyCoverage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//GIVEN
		s := Service{
			TelegramStat: telegramStatMock{get: func(ctx context.Context, query string) (stats *stat.TelegramStats, err error) {
				assert.Equal(t, query, "1129109101", "must be with cut prefix -100")
				return &stat.TelegramStats{ViewsPerPostDaily: 100}, nil
			}},
		}
		//WHEN
		count, err := s.GetDailyCoverage(-1001129109101)
		//THEN
		assert.Nil(t, err)
		assert.Equal(t, 100, count)
	})

	t.Run("error", func(t *testing.T) {
		//GIVEN
		s := Service{
			TelegramStat: telegramStatMock{get: func(ctx context.Context, query string) (stats *stat.TelegramStats, err error) {
				assert.Equal(t, query, "1129109101", "must be with cut prefix -100")
				return nil, fmt.Errorf("went wrong")
			}},
		}
		//WHEN
		count, err := s.GetDailyCoverage(-1001129109101)
		//THEN
		assert.Equal(t, 0, count)
		assert.EqualError(t, err, "went wrong")
	})

	t.Run("channel not found", func(t *testing.T) {
		//GIVEN
		s := Service{
			TelegramStat: telegramStatMock{get: func(ctx context.Context, query string) (stats *stat.TelegramStats, err error) {
				assert.Equal(t, query, "1129109101", "must be with cut prefix -100")
				return nil, stat.ErrChannelNotFound
			}},
		}
		//WHEN
		count, err := s.GetDailyCoverage(-1001129109101)
		//THEN
		assert.Equal(t, 0, count)
		assert.EqualError(t, err, stat.ErrChannelNotFound.Error())
	})
}

type telegramStatMock struct {
	get func(ctx context.Context, query string) (*stat.TelegramStats, error)
}

func (t telegramStatMock) Get(ctx context.Context, query string) (*stat.TelegramStats, error) {
	return t.get(ctx, query)
}
