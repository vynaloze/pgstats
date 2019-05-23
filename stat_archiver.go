package pgstats

import (
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatArchiverView represents content of pg_stat_archiver view
type PgStatArchiverView struct {
	// Number of WAL files that have been successfully archived
	ArchivedCount nullable.Int64 `json:"archived_count"`
	// Name of the last WAL file successfully archived
	LastArchivedWal nullable.String `json:"last_archived_wal"`
	// Time of the last successful archive operation
	LastArchivedTime nullable.Time `json:"last_archived_time"`
	// Number of failed attempts for archiving WAL files
	FailedCount nullable.Int64 `json:"failed_count"`
	// Name of the WAL file of the last failed archival operation
	LastFailedWal nullable.String `json:"last_failed_wal"`
	// Time of the last failed archival operation
	LastFailedTime nullable.Time `json:"last_failed_time"`
	// Time at which these statistics were last reset
	StatsReset nullable.Time `json:"stats_reset"`
}

func (s *PgStats) fetchArchiver() (PgStatArchiverView, error) {
	db := s.conn.db
	query := "select archived_count,last_archived_wal,last_archived_time,failed_count," +
		"last_failed_wal,last_failed_time,stats_reset from pg_stat_archiver"
	row := db.QueryRow(query)
	res := new(PgStatArchiverView)
	err := row.Scan(&res.ArchivedCount, &res.LastArchivedWal, &res.LastArchivedTime, &res.FailedCount,
		&res.LastFailedWal, &res.LastFailedTime, &res.StatsReset)
	return *res, err
}
