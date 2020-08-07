package migrations

func init() {
	include(15, query(`
		create table post (
			id serial primary key,
			lot_id integer not null references lot(id) on delete cascade,
			text text not null,
			buttons jsonb,
			disable_web_page_preview boolean not null,
			scheduled_at timestamp with time zone not null,
			published_at timestamp with time zone
		);

		create index post_lot_id_fkey on post(lot_id);
	`), query(`
			drop table post cascade;
	`))
}
