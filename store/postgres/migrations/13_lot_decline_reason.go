package migrations

func init() {
	include(13, query(`
		alter table lot add column decline_reason text;
	`), query(`
		alter table lot drop column decline_reason;
	`))
}
