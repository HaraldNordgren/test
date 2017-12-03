func (p *PostgreSQL) Hello(v *Type) (int, error) {
	stmt := `
    SELECT
        count(*)
    FROM
        table`

	var count int
	err := p.DB.Get(&count, stmt, v.a)
	return count, err
}
