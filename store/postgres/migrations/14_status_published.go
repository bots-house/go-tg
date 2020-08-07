package migrations

func init() {
	include(14, query(`
       alter type lot_status add value 'scheduled';
    `), query(`

    `))
}
