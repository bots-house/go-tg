package migrations

func init() {
	include(17, query(`
		alter table post add column title text not null;
	`), query(`
		alter table post drop column title;
	`))
}
