package migrations

func init() {
	include(9, query(`
		alter table lot add column views_telegram integer not null default 0;
		alter table lot add column views_site integer not null default 0;
	`), query(`
		alter table lot drop column views_telegram;
		alter table lot drop column views_site;
	`))
}
