package migrations

func init() {
	include(25, query(`
		create table coupon (
			id serial primary key,
			code text not null,
			discount float not null,
			payment_purposes payment_purpose[] not null,
			expire_at timestamp with time zone,
			max_applies_by_user_limit integer,
			max_applies_limit integer,
			is_deleted boolean not null,
			created_at timestamp with time zone not null
		);

		create table coupon_apply (
			id serial primary key,
			coupon_id integer not null references coupon(id) on delete cascade,
			payment_id integer not null references payment(id) on delete cascade,
			unique(coupon_id, payment_id)
		);
		
		create index coupon_apply_coupon_id_fkey on coupon_apply(coupon_id);
		create index coupon_apply_payment_id_fkey on coupon_apply(payment_id);
	`), query(``))
}
