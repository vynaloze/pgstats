package pgstats

import (
	"database/sql"
	"github.com/lib/pq"
)

// PgStatBgWriterView represents content of pg_stat_bgwriter view
type PgStatBgWriterView struct {
	// Number of scheduled checkpoints that have been performed
	CheckpointsTimed sql.NullInt64 `json:"checkpoints_timed"`
	// Number of requested checkpoints that have been performed
	CheckpointsReq sql.NullInt64 `json:"checkpoints_req"`
	// Total amount of time that has been spent in the portion of checkpoint processing
	// where files are written to disk, in milliseconds
	CheckpointWriteTime sql.NullFloat64 `json:"checkpoint_write_time"`
	// Total amount of time that has been spent in the portion of checkpoint processing
	// where files are synchronized to disk, in milliseconds
	CheckpointSyncTime sql.NullFloat64 `json:"checkpoint_sync_time"`
	// Number of buffers written during checkpoints
	BuffersCheckpoint sql.NullInt64 `json:"buffers_checkpoint"`
	// Number of buffers written by the background writer
	BuffersClean sql.NullInt64 `json:"buffers_clean"`
	// Number of times the background writer stopped a cleaning scan because it had written too many buffers
	MaxWrittenClean sql.NullInt64 `json:"maxwritten_clean"`
	// Number of buffers written directly by a backend
	BuffersBackend sql.NullInt64 `json:"buffers_backend"`
	// Number of times a backend had to execute its own fsync call
	// (normally the background writer handles those even when the backend does its own write)
	BuffersBackendFsync sql.NullInt64 `json:"buffers_backend_fsync"`
	// Number of buffers allocated
	BuffersAlloc sql.NullInt64 `json:"buffers_alloc"`
	// Time at which these statistics were last reset
	StatsReset pq.NullTime `json:"stats_reset"`
}

func (s *PgStats) fetchBgWriter() (PgStatBgWriterView, error) {
	db := s.conn.db
	query := "select checkpoints_timed,checkpoints_req,checkpoint_write_time,checkpoint_sync_time,buffers_checkpoint," +
		"buffers_clean,maxwritten_clean,buffers_backend,buffers_backend_fsync,buffers_alloc,stats_reset" +
		" from pg_stat_bgwriter"
	row := db.QueryRow(query)
	res := new(PgStatBgWriterView)
	err := row.Scan(&res.CheckpointsTimed, &res.CheckpointsReq, &res.CheckpointWriteTime, &res.CheckpointSyncTime, &res.BuffersCheckpoint,
		&res.BuffersClean, &res.MaxWrittenClean, &res.BuffersBackend, &res.BuffersBackendFsync, &res.BuffersAlloc, &res.StatsReset)
	return *res, err
}
