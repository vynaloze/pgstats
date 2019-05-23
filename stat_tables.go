package pgstats

import (
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatAllTablesView represents content of pg_stat_all_tables view
type PgStatAllTablesView []PgStatTablesRow

// PgStatSystemTablesView represents content of pg_stat_system_tables view
type PgStatSystemTablesView []PgStatTablesRow

// PgStatUserTablesView represents content of pg_stat_user_tables view
type PgStatUserTablesView []PgStatTablesRow

// PgStatTablesRow represents schema of pg_stat_*_tables views
type PgStatTablesRow struct {
	// OID of a table
	Relid int64 `json:"relid"`
	// Name of the schema that this table is in
	Schemaname string `json:"schemaname"`
	// Name of this table
	Relname string `json:"relname"`
	// Number of sequential scans initiated on this table
	SeqScan nullable.Int64 `json:"seq_scan"`
	// Number of live rows fetched by sequential scans
	SeqTupRead nullable.Int64 `json:"seq_tup_read"`
	// Number of index scans initiated on this table
	IdxScan nullable.Int64 `json:"idx_scan"`
	// Number of live rows fetched by index scans
	IdxTupFetch nullable.Int64 `json:"idx_tup_fetch"`
	// Number of rows inserted
	NTupIns nullable.Int64 `json:"n_tup_ins"`
	// Number of rows updated (includes HOT updated rows)
	NTupUpd nullable.Int64 `json:"n_tup_upd"`
	// Number of rows deleted
	NTupDel nullable.Int64 `json:"n_tup_del"`
	// Number of rows HOT updated (i.e., with no separate index update required)
	NTupHotUpd nullable.Int64 `json:"n_tup_hot_upd"`
	// Estimated number of live rows
	NLiveTup nullable.Int64 `json:"n_live_tup"`
	// Estimated number of dead rows
	NDeadTup nullable.Int64 `json:"n_dead_tup"`
	// Estimated number of rows modified since this table was last analyzed
	NModSinceAnalyze nullable.Int64 `json:"n_mod_since_analyze"`
	// Last time at which this table was manually vacuumed (not counting VACUUM FULL)
	LastVacuum nullable.Time `json:"last_vacuum"`
	// Last time at which this table was vacuumed by the autovacuum daemon
	LastAutovacuum nullable.Time `json:"last_autovacuum"`
	// Last time at which this table was manually analyzed
	LastAnalyze nullable.Time `json:"last_analyze"`
	// Last time at which this table was analyzed by the autovacuum daemon
	LastAutoanalyze nullable.Time `json:"last_autoanalyze"`
	// Number of times this table has been manually vacuumed (not counting VACUUM FULL)
	VacuumCount nullable.Int64 `json:"vacuum_count"`
	// Number of times this table has been vacuumed by the autovacuum daemon
	AutovacuumCount nullable.Int64 `json:"autovacuum_count"`
	// Number of times this table has been manually analyzed
	AnalyzeCount nullable.Int64 `json:"analyze_count"`
	// Number of times this table has been analyzed by the autovacuum daemon
	AutoanalyzeCount nullable.Int64 `json:"autoanalyze_count"`
}

func (s *PgStats) fetchTables(view string) ([]PgStatTablesRow, error) {
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

	data := make([]PgStatTablesRow, 0)
	for rows.Next() {
		row := new(PgStatTablesRow)
		err := rows.Scan(&row.Relid, &row.Schemaname, &row.Relname, &row.SeqScan, &row.SeqTupRead,
			&row.IdxScan, &row.IdxTupFetch, &row.NTupIns, &row.NTupUpd, &row.NTupDel,
			&row.NTupHotUpd, &row.NLiveTup, &row.NDeadTup, &row.NModSinceAnalyze, &row.LastVacuum,
			&row.LastAutovacuum, &row.LastAnalyze, &row.LastAutoanalyze, &row.VacuumCount, &row.AutovacuumCount,
			&row.AnalyzeCount, &row.AutoanalyzeCount)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
