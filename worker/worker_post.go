package worker

import "context"

func (wrk *Worker) taskPublishPosts(ctx context.Context) error {
	return wrk.Posting.SendPosts(ctx)
}
