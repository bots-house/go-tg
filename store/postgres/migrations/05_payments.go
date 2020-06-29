package migrations

func init() {
	include(5, query(`
		create type lot_status as enum (
			'created',
			'paid',
			'published',
			'declined',
			'canceled'
		);

		create table lot_canceled_reason (
			id serial primary key, 
			why text not null, 
			is_public boolean not null, 
			created_at timestamp with time zone not null,
			updated_at timestamp with time zone
		);

		-- create lot status field
		alter table lot
			add column status lot_status not null default('created');
		alter table lot
			alter column status drop default;

		-- create lot -> canceled_reason relation;
		alter table lot
			add column canceled_reason_id integer 
			    references lot_canceled_reason(id) on delete set null;
        
        create index lot_status_idx on lot(status);
        
        create type payment_status as enum (
              'created',
              'pending',
              'success',
              'failed'
        );

        create type payment_purpose as enum (
            'application',
            'change_price'
        );

        create type payment_gateway as enum (
            'interkassa',
            'direct'
        );

        create table payment (
              id serial primary key,
              external_id text,
              gateway payment_gateway not null, 
              purpose payment_purpose not null,
              payer_id integer not null references "user"(id) on delete set null,
              lot_id integer not null references "lot"(id) on delete set null,
              status payment_status not null, 
              requested jsonb not null, 
              paid jsonb, 
              received jsonb, 
              metadata jsonb,
              created_at timestamp with time zone not null,
              updated_at timestamp with time zone
        );
	`), query(`
        delete type lot_status;
        drop table lot_canceled_reason;
        delete type payment_status;
        delete type payment_purpose;
        delete type payment_gateway;
        drop table payment;
    `))
}
