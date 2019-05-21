package pgstats

// PgStatStatementsView represents content of pg_stat_statements view
type PgStatStatementsView []PgStatStatementsRow

// PgStatStatementsRow represents schema of pg_stat_statements view
type PgStatStatementsRow struct {
	// OID of user who executed the statement
	Userid int64 `json:"userid"`
	// OID of database in which the statement was executed
	Dbid int64 `json:"dbid"`
	// Internal hash code, computed from the statement's parse tree
	Queryid int64 `json:"queryid"`
	// Text of a representative statement
	Query string `json:"query"`
	// Number of times executed
	Calls int64 `json:"calls"`
	// Total time spent in the statement, in milliseconds
	TotalTime float64 `json:"total_time"`
	// Minimum time spent in the statement, in milliseconds
	MinTime float64 `json:"min_time"`
	// Maximum time spent in the statement, in milliseconds
	MaxTime float64 `json:"max_time"`
	// Mean time spent in the statement, in milliseconds
	MeanTime float64 `json:"mean_time"`
	// Population standard deviation of time spent in the statement, in milliseconds
	StddevTime float64 `json:"stddev_time"`
	// Total number of rows retrieved or affected by the statement
	Rows int64 `json:"rows"`
	// Total number of shared block cache hits by the statement
	SharedBlksHit int64 `json:"shared_blks_hit"`
	// Total number of shared blocks read by the statement
	SharedBlksRead int64 `json:"shared_blks_read"`
	// Total number of shared blocks dirtied by the statement
	SharedBlksDirtied int64 `json:"shared_blks_dirtied"`
	// Total number of shared blocks written by the statement
	SharedBlksWritten int64 `json:"shared_blks_written"`
	// Total number of local block cache hits by the statement
	LocalBlksHit int64 `json:"local_blks_hit"`
	// Total number of local blocks read by the statement
	LocalBlksRead int64 `json:"local_blks_read"`
	// Total number of local blocks dirtied by the statement
	LocalBlksDirtied int64 `json:"local_blks_dirtied"`
	// Total number of local blocks written by the statement
	LocalBlksWritten int64 `json:"local_blks_written"`
	// Total number of temp blocks read by the statement
	TempBlksRead int64 `json:"temp_blks_read"`
	// Total number of temp blocks written by the statement
	TempBlksWritten int64 `json:"temp_blks_written"`
	// Total time the statement spent reading blocks, in milliseconds (if track_io_timing is enabled, otherwise zero)
	BlkReadTime float64 `json:"blk_read_time"`
	// Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero)
	BlkWriteTime float64 `json:"blk_write_time"`
}

func (s *PgStats) fetchStatements() (PgStatStatementsView, error) {
	db := s.conn.db
	query := "select userid,dbid,queryid,query,calls," +
		"total_time,min_time,max_time,mean_time,stddev_time," +
		"rows,shared_blks_hit,shared_blks_read,shared_blks_dirtied,shared_blks_written," +
		"local_blks_hit,local_blks_read,local_blks_dirtied,local_blks_written,temp_blks_read," +
		"temp_blks_written,blk_read_time,blk_write_time from pg_stat_statements"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatStatementsRow, 0)
	for rows.Next() {
		row := new(PgStatStatementsRow)
		err := rows.Scan(&row.Userid, &row.Dbid, &row.Queryid, &row.Query, &row.Calls,
			&row.TotalTime, &row.MinTime, &row.MaxTime, &row.MeanTime, &row.StddevTime,
			&row.Rows, &row.SharedBlksHit, &row.SharedBlksRead, &row.SharedBlksDirtied, &row.SharedBlksWritten,
			&row.LocalBlksHit, &row.LocalBlksRead, &row.LocalBlksDirtied, &row.LocalBlksWritten, &row.TempBlksRead,
			&row.TempBlksWritten, &row.BlkReadTime, &row.BlkWriteTime)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
