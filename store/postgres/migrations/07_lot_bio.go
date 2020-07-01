package migrations

func init() {
	include(7, query(`
		alter table lot add column bio text;
	`), query(`
        alter table drop column bio;
    `))
}
