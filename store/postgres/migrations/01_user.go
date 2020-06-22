package migrations

func init() {
	include(1, query(`
        create type joined_from as enum ('site', 'bot');

        create table "user" (
            id serial primary key,

            telegram_id integer unique not null,
            telegram_username varchar(32),
            telegram_language_code text,

            first_name text not null,
            last_name text,
            is_name_edited bool not null,

            avatar text,
            is_admin bool not null,

            joined_from joined_from,

            joined_at timestamp with time zone not null,
            updated_at timestamp with time zone
        );


    `), query(`
        drop table "user";
        drop type joined_from;

    `))
}
