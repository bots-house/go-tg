package migrations

func init() {
	include(22, query(`
		update lot set metrics_monthly_income = 0 where metrics_monthly_income is null;
		alter table lot alter column metrics_monthly_income set not null;
	`), query(``))
}
