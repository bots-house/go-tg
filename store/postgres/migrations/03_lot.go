package migrations

func init() {
	include(3, query(`
        create table lot (
            id serial primary key,
  			owner_id integer not null references "user"(id) on delete cascade,
			external_id bigint not null,
			name text not null,
			avatar text,
			username text, 
			join_link text,
            price_current int not null, 
			price_previous int,
        	price_is_bargain bool not null, 
			comment text not null, 
			metrics_members_count int not null, 
			metrics_daily_coverage int not null, 
			metrics_monthly_income int,
			metrics_price_per_member double precision not null, 
			metrics_price_per_view double precision not null, 
			metrics_payback_period double precision,
			extra_resources jsonb, 
			created_at timestamp with time zone not null,
			paid_at timestamp with time zone, 
			approved_at timestamp with time zone,
			published_at timestamp with time zone
		);	
		
		create table lot_topic (
			id serial primary key, 
			lot_id integer not null references lot(id) on delete cascade,
			topic_id integer not null references topic(id) on delete cascade,
			unique(lot_id, topic_id)
		);

		create index topic_lot_lot_id_idx ON lot_topic(lot_id);
    `), query(`
		drop table lot_topic;
        drop table lot;
    `))
}
