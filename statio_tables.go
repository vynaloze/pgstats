package pgstats

import (
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatIoAllTablesView represents content of pg_statio_all_tables view
type PgStatIoAllTablesView []PgStatIoTablesRow

// PgStatIoSystemTablesView represents content of pg_statio_system_tables view
type PgStatIoSystemTablesView []PgStatIoTablesRow

// PgStatIoUserTablesView represents content of pg_statio_user_tables view
type PgStatIoUserTablesView []PgStatIoTablesRow

// PgStatIoTablesRow represents schema of pg_statio_*_tables views
type PgStatIoTablesRow struct {
	// OID of a table
	Relid int64 `json:"relid"`
	// Name of the schema that this table is in
	Schemaname string `json:"schemaname"`
	// Name of this table
	Relname string `json:"relname"`
	// Number of disk blocks read from this table
	HeapBlksRead nullable.Int64 `json:"heap_blks_read"`
	// Number of buffer hits in this table
	HeapBlksHit nullable.Int64 `json:"heap_blks_hit"`
	// Number of disk blocks read from all indexes on this table
	IdxBlksRead nullable.Int64 `json:"idx_blks_read"`
	// Number of buffer hits in all indexes on this table
	IdxBlksHit nullable.Int64 `json:"idx_blks_hit"`
	// Number of disk blocks read from this table's TOAST table (if any)
	ToastBlksRead nullable.Int64 `json:"toast_blks_read"`
	// Number of buffer hits in this table's TOAST table (if any)
	ToastBlksHit nullable.Int64 `json:"toast_blks_hit"`
	// Number of disk blocks read from this table's TOAST table indexes (if any)
	TidxBlksRead nullable.Int64 `json:"tidx_blks_read"`
	// Number of buffer hits in this table's TOAST table indexes (if any)
	TidxBlksHit nullable.Int64 `json:"tidx_blks_hit"`
}

func (s *PgStats) fetchIoTables(view string) ([]PgStatIoTablesRow, error) {
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

	data := make([]PgStatIoTablesRow, 0)
	for rows.Next() {
		row := new(PgStatIoTablesRow)
		err := rows.Scan(&row.Relid, &row.Schemaname, &row.Relname,
			&row.HeapBlksRead, &row.HeapBlksHit,
			&row.IdxBlksRead, &row.IdxBlksHit,
			&row.ToastBlksRead, &row.ToastBlksHit,
			&row.TidxBlksRead, &row.TidxBlksHit)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
