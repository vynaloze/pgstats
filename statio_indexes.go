package pgstats

import "database/sql"

// PgStatIoAllIndexesView represents content of pg_statio_all_indexes view
type PgStatIoAllIndexesView []PgStatIoIndexesRow

// PgStatIoSystemIndexesView represents content of pg_statio_system_indexes view
type PgStatIoSystemIndexesView []PgStatIoIndexesRow

// PgStatIoUserIndexesView represents content of pg_statio_user_indexes view
type PgStatIoUserIndexesView []PgStatIoIndexesRow

// PgStatIoIndexesRow represents schema of pg_statio_*_indexes views
type PgStatIoIndexesRow struct {
	// OID of the table for this index
	RelId int64
	// OID of this index
	IndexRelId int64
	// Name of the schema this index is in
	SchemaName string
	// Name of the table for this index
	RelName string
	// Name of this index
	IndexRelName string
	// Number of disk blocks read from this index
	IdxBlksRead sql.NullInt64
	// Number of buffer hits in this index
	IdxBlksHit sql.NullInt64
}

func (s *PgStats) fetchIoIndexes(view string) ([]PgStatIoIndexesRow, error) {
	db := s.conn.db
	query := "select relid,indexrelid,schemaname,relname,indexrelname," +
		"idx_blks_read,idx_blks_hit from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatIoIndexesRow, 0)
	for rows.Next() {
		row := new(PgStatIoIndexesRow)
		err := rows.Scan(&row.RelId, &row.IndexRelId, &row.SchemaName, &row.RelName, &row.IndexRelName,
			&row.IdxBlksRead, &row.IdxBlksHit)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
