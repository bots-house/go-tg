package migrations

func init() {
	include(8, query(`
		alter table review alter column telegram_id drop not null;
	`), query(`
		alter table review drop column telegram_id;
	`))
}
