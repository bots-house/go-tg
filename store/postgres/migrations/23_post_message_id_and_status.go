package migrations

func init() {
	include(23, query(`
		alter table post add column message_id integer;
		create type post_status as enum ('scheduled', 'published');
		alter table post add column status post_status not null;
	`), query(`
		alter table post drop column message_id;
		alter table post drop column status;
		drop type post_status;
	`))
}
