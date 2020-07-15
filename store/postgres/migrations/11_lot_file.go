package migrations

func init() {
	include(11, query(`
		create table lot_file (
			id serial primary key,
			lot_id integer references lot(id) on delete cascade,
			name text not null,
			size integer not null,
			mime_type text not null,
			path text not null,
			created_at timestamp with time zone not null
		);

		create index lot_file_lot_id_fkey on lot_file(lot_id);
	`), query(`
		drop table lot_file;
	`))
}
