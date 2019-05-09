package pgstats

import (
	"database/sql"
)

// PgStatXactAllTablesView represents content of pg_stat_xact_all_tables view
type PgStatXactAllTablesView []PgStatXactTablesRow

// PgStatXactSystemTablesView represents content of pg_stat_xact_system_tables view
type PgStatXactSystemTablesView []PgStatXactTablesRow

// PgStatXactUserTablesView represents content of pg_stat_xact_user_tables view
type PgStatXactUserTablesView []PgStatXactTablesRow

// PgStatXactTablesRow represents schema of pg_stat_xact_*_tables views
type PgStatXactTablesRow struct {
	// OID of a table
	RelId int64
	// Name of the schema that this table is in
	SchemaName string
	// Name of this table
	RelName string
	// Number of sequential scans initiated on this table
	SeqScan sql.NullInt64
	// Number of live rows fetched by sequential scans
	SeqTupRead sql.NullInt64
	// Number of index scans initiated on this table
	IdxScan sql.NullInt64
	// Number of live rows fetched by index scans
	IdxTupFetch sql.NullInt64
	// Number of rows inserted
	NTupIns sql.NullInt64
	// Number of rows updated (includes HOT updated rows)
	NTupUpd sql.NullInt64
	// Number of rows deleted
	NTupDel sql.NullInt64
	// Number of rows HOT updated (i.e., with no separate index update required)
	NTupHotUpd sql.NullInt64
}

func (s *PgStats) fetchXactTables(view string) ([]PgStatXactTablesRow, error) {
	data := make([]PgStatXactTablesRow, 0)
	db := s.conn.db
	query := "select relid,schemaname,relname,seq_scan,seq_tup_read," +
		"idx_scan,idx_tup_fetch,n_tup_ins,n_tup_upd,n_tup_del,n_tup_hot_upd from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(PgStatXactTablesRow)
		err := rows.Scan(&row.RelId, &row.SchemaName, &row.RelName, &row.SeqScan, &row.SeqTupRead,
			&row.IdxScan, &row.IdxTupFetch, &row.NTupIns, &row.NTupUpd, &row.NTupDel, &row.NTupHotUpd)
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
