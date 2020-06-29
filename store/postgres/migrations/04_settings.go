package migrations

func init() {
	include(4, query(`
		create table settings (
			id integer primary key,

			prices_application jsonb not null,
			prices_change jsonb not null, 
			
			channel_public_username text not null,
			channel_private_link text not null,
			channel_private_id bigint not null, 

			cashier_username text not null,

			updated_at timestamp with time zone
		);
    `), query(`
		drop table settings;
    `))
}
