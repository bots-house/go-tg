package migrations

func init() {
	include(10, query(`
		create table favorite (
			id serial primary key,
			lot_id integer not null references lot(id) on delete cascade,
			user_id integer not null references "user"(id) on delete cascade,
			created_at timestamp with time zone not null
		);

		create index favorite_lot_id_fkey on favorite(lot_id);
		create index favorite_user_id_fkey on favorite(user_id);
	`), query(`
		drop table favorite;
	`))
}
