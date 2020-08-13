package migrations

func init() {
	include(20, query(`
		alter table post alter column lot_id drop not null;
		alter table post alter column title drop not null;
	`), query(`
	`))
}
