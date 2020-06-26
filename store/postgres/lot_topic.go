package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/bots-house/birzzha/core"
	"github.com/bots-house/birzzha/store"
	"github.com/bots-house/birzzha/store/postgres/shared"
)

type LotTopicStore struct {
	db *sql.DB
	txier store.Txier
}

func (ls *LotTopicStore) set(ctx context.Context, lot core.LotID, topics []core.TopicID) error {
	tx := shared.GetTx(ctx)

	stmt, err := tx.PrepareContext(ctx, `insert into lot_topic(lot_id, topic_id) values ($1, $2)`)
	if err != nil {
		return errors.Wrap(err, "prepare stmt")
	}
	defer stmt.Close()

	for i, id := range topics {
		_, err := stmt.ExecContext(ctx, lot, id)
		if err != nil {
			return errors.Wrapf(err, "insert lot_topic #%d", i)
		}
	}

	return nil
}

func (ls *LotTopicStore) delete(ctx context.Context, lot core.LotID) error {
	tx := shared.GetTx(ctx)

	_, err := tx.ExecContext(ctx, `delete from lot_topic where lot_id = $1`, lot)
	if err != nil {
		return errors.Wrap(err, "exec delete lot topic")
	}

	return nil
}


func (ls *LotTopicStore) Set(ctx context.Context, lot core.LotID, topics []core.TopicID) error {
	return ls.txier(ctx, func(ctx context.Context) error {
		if err := ls.delete(ctx, lot); err != nil {
			return errors.Wrap(err, "delete existing rel")
		}

		if err := ls.set(ctx, lot, topics); err != nil {
			return errors.Wrap(err, "create new rel")
		}

		return nil
	})
}

func (ls *LotTopicStore) Get(ctx context.Context, lotID core.LotID) (core.TopicSlice, error) {
	executor := shared.GetExecutorOrDefault(ctx, ls.db)

	rows, err := executor.QueryContext(ctx, `
		select topic.id, topic.name, topic.slug, topic.created_at
		from lot_topic
		inner join topic on lot_topic.topic_id = topic.id
		where lot_id = $1
	`, lotID)
	if err != nil {
		return nil, errors.Wrap(err, "query rows")
	}
	defer rows.Close()

	result := make(core.TopicSlice, 0, 3)

	for rows.Next() {
		item := &core.Topic{}

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Slug,
			&item.CreatedAt,
		); err != nil {
			return nil, errors.Wrap(err, "scan")
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows err")
	}

	return result, nil
}
