package pgstats

import (
	"database/sql"
)

// PgStatIoAllTablesView represents content of pg_statio_all_tables view
type PgStatIoAllTablesView []PgStatIoTablesRow

// PgStatIoUserTablesView represents content of pg_statio_user_tables view
type PgStatIoUserTablesView []PgStatIoTablesRow

// PgStatIoSystemTablesView represents content of pg_statio_system_tables view
type PgStatIoSystemTablesView []PgStatIoTablesRow

// PgStatIoTablesRow represents schema of pg_statio_*_tables views
type PgStatIoTablesRow struct {
	// OID of a table
	RelId int64
	// Name of the schema that this table is in
	SchemaName string
	// Name of this table
	RelName string
	// Number of disk blocks read from this table
	HeapBlksRead sql.NullInt64
	// Number of buffer hits in this table
	HeapBlksHit sql.NullInt64
	// Number of disk blocks read from all indexes on this table
	IdxBlksRead sql.NullInt64
	// Number of buffer hits in all indexes on this table
	IdxBlksHit sql.NullInt64
	// Number of disk blocks read from this table's TOAST table (if any)
	ToastBlksRead sql.NullInt64
	// Number of buffer hits in this table's TOAST table (if any)
	ToastBlksHit sql.NullInt64
	// Number of disk blocks read from this table's TOAST table indexes (if any)
	TidxBlksRead sql.NullInt64
	// Number of buffer hits in this table's TOAST table indexes (if any)
	TidxBlksHit sql.NullInt64
}

func (s *PgStats) fetchIoTables(view string) ([]PgStatIoTablesRow, error) {
	data := make([]PgStatIoTablesRow, 0)
	db := s.conn.db
	query := "select relid,schemaname,relname," +
		"heap_blks_read,heap_blks_hit," +
		"idx_blks_read,idx_blks_hit," +
		"toast_blks_read,toast_blks_hit," +
		"tidx_blks_read,tidx_blks_hit from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(PgStatIoTablesRow)
		err := rows.Scan(&row.RelId, &row.SchemaName, &row.RelName,
			&row.HeapBlksRead, &row.HeapBlksHit,
			&row.IdxBlksRead, &row.IdxBlksHit,
			&row.ToastBlksRead, &row.ToastBlksHit,
			&row.TidxBlksRead, &row.TidxBlksHit)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
