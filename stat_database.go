package pgstats

import (
	"database/sql"
	"github.com/lib/pq"
)

// PgStatDatabaseView represents content of pg_stat_database view
type PgStatDatabaseView []PgStatDatabaseRow

// PgStatDatabaseRow represents schema of pg_stat_database view
type PgStatDatabaseRow struct {
	// OID of a database
	DatId int64
	// Name of this database
	DatName string
	// Number of backends currently connected to this database.
	// This is the only column in this view that returns a value reflecting current state;
	// all other columns return the accumulated values since the last reset.
	NumBackends int64
	// Number of transactions in this database that have been committed
	XactCommit sql.NullInt64
	//	Number of transactions in this database that have been rolled back
	XactRollback sql.NullInt64
	// Number of disk blocks read in this database
	BlksRead sql.NullInt64
	// Number of times disk blocks were found already in the buffer cache, so that a read was not necessary
	// (this only includes hits in the PostgreSQL buffer cache, not the operating system's file system cache)
	BlksHit sql.NullInt64
	// Number of rows returned by queries in this database
	TupReturned sql.NullInt64
	// Number of rows fetched by queries in this database
	TupFetched sql.NullInt64
	// Number of rows inserted by queries in this database
	TupInserted sql.NullInt64
	// Number of rows updated by queries in this database
	TupUpdated sql.NullInt64
	// 	Number of rows deleted by queries in this database
	TupDeleted sql.NullInt64
	// Number of queries canceled due to conflicts with recovery in this database.
	// (Conflicts occur only on standby servers; see pg_stat_database_conflicts for details.)
	Conflicts sql.NullInt64
	// Number of temporary files created by queries in this database.
	// All temporary files are counted,
	// regardless of why the temporary file was created (e.g., sorting or hashing),
	// and regardless of the log_temp_files setting.
	TempFiles sql.NullInt64
	// Total amount of data written to temporary files by queries in this database.
	// All temporary files are counted,
	// regardless of why the temporary file was created,
	// and regardless of the log_temp_files setting.
	TempBytes sql.NullInt64
	// Number of deadlocks detected in this database
	Deadlocks sql.NullInt64
	// Time spent reading data file blocks by backends in this database, in milliseconds
	BlkReadTime sql.NullFloat64
	// Time spent writing data file blocks by backends in this database, in milliseconds
	BlkWriteTime sql.NullFloat64
	// Time at which these statistics were last reset
	StatsReset pq.NullTime
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
		err := rows.Scan(&row.DatId, &row.DatName, &row.NumBackends, &row.XactCommit, &row.XactRollback,
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
