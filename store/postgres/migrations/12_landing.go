package migrations

func init() {
	include(12, query(`
        create table landing (
            id integer primary key,

            unique_users_per_month_actual integer not null,
            unique_users_per_month_shift integer not null,

            avg_site_reach_actual integer not null,
            avg_site_reach_shift integer not null,

            avg_channel_reach_actual integer not null,
            avg_channel_reach_shift integer not null
        )
    `), query(`
		drop table landing;
    `))
}
