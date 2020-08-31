package migrations

func init() {
	include(21, query(`
		alter type payment_gateway add value 'unitpay';
		alter type payment_status add value 'error';
		alter table payment add constraint unique_external_id_gateway unique(external_id, gateway);
	`), query(`
		alter table payment drop constraint unique_external_id_gateway;
	`))
}
