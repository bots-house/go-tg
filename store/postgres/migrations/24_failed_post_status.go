package migrations

func init() {
	include(24, query(`
		alter type post_status add value 'failed' after 'published';
	`), query(`
	`))
}
