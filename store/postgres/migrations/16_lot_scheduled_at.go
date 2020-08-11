package migrations

func init() {
	include(16, query(`
		alter table lot add column scheduled_at timestamp with time zone;
	`), query(`
		alter table lot drop column scheduled_at;
	`))
}
