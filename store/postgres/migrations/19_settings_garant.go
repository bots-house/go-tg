package migrations

func init() {
	include(19, query(`
		alter table settings add column garant_name text not null default 'Alex Dalakian';
		alter table settings add column garant_username text not null default 'alexxdd'; 
		alter table settings add column garant_reviews_channel text not null default 'birzzha_review';
		alter table settings add column garant_percentage_deal double precision not null default 2;
		alter table settings add column garant_percentage_deal_discount_period double precision;
		alter table settings add column garant_avatar_url text;
	`), query(`
		alter table settings drop column garant_name;
		alter table settings drop column garant_username;
		alter table settings drop column garant_reviews_channel;
		alter table settings drop column garant_percentage_deal;
		alter table settings drop column garant_percentage_deal_discount_period;
		alter table settings drop column garant_avatar_url;
	`))
}
