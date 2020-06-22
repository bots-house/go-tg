package migrations

func init() {
	include(2, query(`
        create table topic (
            id serial primary key,
            name text not null,
            slug text unique not null,
            created_at timestamp with time zone not null
        );


    `), query(`
        drop table topic;
    `))
}
