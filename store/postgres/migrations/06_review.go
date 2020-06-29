package migrations

func init() {
	include(6, query(`
		create table review (
			id serial primary key,
			telegram_id integer not null,
			first_name text not null,
			last_name text,
			username varchar(32),
			avatar text,
			text text not null,
			created_at timestamp with time zone not null
		);
	`), query(`
			drop table review;
	`))
}
