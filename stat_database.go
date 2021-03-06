package pgstats

import (
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatDatabaseView represents content of pg_stat_database view
type PgStatDatabaseView []PgStatDatabaseRow

// PgStatDatabaseRow represents schema of pg_stat_database view
type PgStatDatabaseRow struct {
	// OID of a database
	Datid int64 `json:"datid"`
	// Name of this database
	Datname string `json:"datname"`
	// Number of backends currently connected to this database.
	// This is the only column in this view that returns a value reflecting current state;
	// all other columns return the accumulated values since the last reset.
	NumBackends int64 `json:"numbackends"`
	// Number of transactions in this database that have been committed
	XactCommit nullable.Int64 `json:"xact_commit"`
	//	Number of transactions in this database that have been rolled back
	XactRollback nullable.Int64 `json:"xact_rollback"`
	// Number of disk blocks read in this database
	BlksRead nullable.Int64 `json:"blks_read"`
	// Number of times disk blocks were found already in the buffer cache, so that a read was not necessary
	// (this only includes hits in the PostgreSQL buffer cache, not the operating system's file system cache)
	BlksHit nullable.Int64 `json:"blks_hit"`
	// Number of rows returned by queries in this database
	TupReturned nullable.Int64 `json:"tup_returned"`
	// Number of rows fetched by queries in this database
	TupFetched nullable.Int64 `json:"tup_fetched"`
	// Number of rows inserted by queries in this database
	TupInserted nullable.Int64 `json:"tup_inserted"`
	// Number of rows updated by queries in this database
	TupUpdated nullable.Int64 `json:"tup_updated"`
	// 	Number of rows deleted by queries in this database
	TupDeleted nullable.Int64 `json:"tup_deleted"`
	// Number of queries canceled due to conflicts with recovery in this database.
	// (Conflicts occur only on standby servers; see pg_stat_database_conflicts for details.)
	Conflicts nullable.Int64 `json:"conflicts"`
	// Number of temporary files created by queries in this database.
	// All temporary files are counted,
	// regardless of why the temporary file was created (e.g., sorting or hashing),
	// and regardless of the log_temp_files setting.
	TempFiles nullable.Int64 `json:"temp_files"`
	// Total amount of data written to temporary files by queries in this database.
	// All temporary files are counted,
	// regardless of why the temporary file was created,
	// and regardless of the log_temp_files setting.
	TempBytes nullable.Int64 `json:"temp_bytes"`
	// Number of deadlocks detected in this database
	Deadlocks nullable.Int64 `json:"deadlocks"`
	// Time spent reading data file blocks by backends in this database, in milliseconds
	BlkReadTime nullable.Float64 `json:"blk_read_time"`
	// Time spent writing data file blocks by backends in this database, in milliseconds
	BlkWriteTime nullable.Float64 `json:"blk_write_time"`
	// Time at which these statistics were last reset
	StatsReset nullable.Time `json:"stats_reset"`
}

func (s *PgStats) fetchDatabases() ([]PgStatDatabaseRow, error) {
	db := s.conn.db
	query := "select datid,datname,numbackends,xact_commit,xact_rollback," +
		"blks_read,blks_hit,tup_returned,tup_fetched,tup_inserted," +
		"tup_updated,tup_deleted,conflicts,temp_files,temp_bytes," +
		"deadlocks,blk_read_time,blk_write_time,stats_reset from pg_stat_database"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatDatabaseRow, 0)
	for rows.Next() {
		row := new(PgStatDatabaseRow)
		err := rows.Scan(&row.Datid, &row.Datname, &row.NumBackends, &row.XactCommit, &row.XactRollback,
			&row.BlksRead, &row.BlksHit, &row.TupReturned, &row.TupFetched, &row.TupInserted,
			&row.TupUpdated, &row.TupDeleted, &row.Conflicts, &row.TempFiles, &row.TempBytes,
			&row.Deadlocks, &row.BlkReadTime, &row.BlkWriteTime, &row.StatsReset)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
