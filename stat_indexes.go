package pgstats

import "database/sql"

// PgStatAllIndexesView represents content of pg_stat_all_indexes view
type PgStatAllIndexesView []PgStatIndexesRow

// PgStatSystemIndexesView represents content of pg_stat_system_indexes view
type PgStatSystemIndexesView []PgStatIndexesRow

// PgStatUserIndexesView represents content of pg_stat_user_indexes view
type PgStatUserIndexesView []PgStatIndexesRow

// PgStatIndexesRow represents schema of pg_stat_*_indexes views
type PgStatIndexesRow struct {
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
	// Number of index scans initiated on this index
	IdxScan sql.NullInt64
	// Number of index entries returned by scans on this index
	IdxTupRead sql.NullInt64
	// Number of live table rows fetched by simple index scans using this index
	IdxTupFetch sql.NullInt64
}

func (s *PgStats) fetchIndexes(view string) ([]PgStatIndexesRow, error) {
	db := s.conn.db
	query := "select relid,indexrelid,schemaname,relname,indexrelname," +
		"idx_scan,idx_tup_read,idx_tup_fetch from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatIndexesRow, 0)
	for rows.Next() {
		row := new(PgStatIndexesRow)
		err := rows.Scan(&row.RelId, &row.IndexRelId, &row.SchemaName, &row.RelName, &row.IndexRelName,
			&row.IdxScan, &row.IdxTupRead, &row.IdxTupFetch)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
