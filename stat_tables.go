package pgstats

import (
	"database/sql"
	"github.com/lib/pq"
)

// PgStatAllTablesView represents content of pg_stat_all_tables view
type PgStatAllTablesView []PgStatTablesRow

// PgStatUserTablesView represents content of pg_stat_user_tables view
type PgStatUserTablesView []PgStatTablesRow

// PgStatSystemTablesView represents content of pg_stat_system_tables view
type PgStatSystemTablesView []PgStatTablesRow

// PgStatTablesRow represents schema of pg_stat_*_tables views
type PgStatTablesRow struct {
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
	// Estimated number of live rows
	NLiveTup sql.NullInt64
	// Estimated number of dead rows
	NDeadTup sql.NullInt64
	// Estimated number of rows modified since this table was last analyzed
	NModSinceAnalyze sql.NullInt64
	// Last time at which this table was manually vacuumed (not counting VACUUM FULL)
	LastVacuum pq.NullTime
	// Last time at which this table was vacuumed by the autovacuum daemon
	LastAutovacuum pq.NullTime
	// Last time at which this table was manually analyzed
	LastAnalyze pq.NullTime
	// Last time at which this table was analyzed by the autovacuum daemon
	LastAutoanalyze pq.NullTime
	// Number of times this table has been manually vacuumed (not counting VACUUM FULL)
	VacuumCount sql.NullInt64
	// Number of times this table has been vacuumed by the autovacuum daemon
	AutovacuumCount sql.NullInt64
	// Number of times this table has been manually analyzed
	AnalyzeCount sql.NullInt64
	// Number of times this table has been analyzed by the autovacuum daemon
	AutoanalyzeCount sql.NullInt64
}

func (s *PgStats) fetchTables(view string) ([]PgStatTablesRow, error) {
	data := make([]PgStatTablesRow, 0)
	db := s.conn.db
	query := "select relid,schemaname,relname,seq_scan,seq_tup_read," +
		"idx_scan,idx_tup_fetch,n_tup_ins,n_tup_upd,n_tup_del," +
		"n_tup_hot_upd,n_live_tup,n_dead_tup,n_mod_since_analyze,last_vacuum," +
		"last_autovacuum,last_analyze,last_autoanalyze,vacuum_count,autovacuum_count," +
		"analyze_count,autoanalyze_count from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(PgStatTablesRow)
		err := rows.Scan(&row.RelId, &row.SchemaName, &row.RelName, &row.SeqScan, &row.SeqTupRead,
			&row.IdxScan, &row.IdxTupFetch, &row.NTupIns, &row.NTupUpd, &row.NTupDel,
			&row.NTupHotUpd, &row.NLiveTup, &row.NDeadTup, &row.NModSinceAnalyze, &row.LastVacuum,
			&row.LastAutovacuum, &row.LastAnalyze, &row.LastAutoanalyze, &row.VacuumCount, &row.AutovacuumCount,
			&row.AnalyzeCount, &row.AutoanalyzeCount)
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
