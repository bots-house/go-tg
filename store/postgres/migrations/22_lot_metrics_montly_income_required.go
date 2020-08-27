package migrations

func init() {
	include(22, query(`
		alter table lot alter column metrics_monthly_income set default 0;
		alter table lot alter column metrics_monthly_income set not null;
		alter table lot alter column metrics_monthly_income drop default;
	`), query(``))
}
