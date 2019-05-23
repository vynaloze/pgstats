package pgstats

import (
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatDatabaseConflictsView represents content of pg_stat_database_conflicts view
type PgStatDatabaseConflictsView []PgStatDatabaseConflictsRow

// PgStatDatabaseConflictsRow represents schema of pg_stat_database_conflicts view
type PgStatDatabaseConflictsRow struct {
	// OID of a database
	Datid int64 `json:"datid"`
	// Name of this database
	Datname string `json:"datname"`
	// Number of queries in this database that have been canceled due to dropped tablespaces
	ConflTablespace nullable.Int64 `json:"confl_tablespace"`
	// Number of queries in this database that have been canceled due to lock timeouts
	ConflLock nullable.Int64 `json:"confl_lock"`
	// Number of queries in this database that have been canceled due to old snapshots
	ConflSnapshot nullable.Int64 `json:"confl_snapshot"`
	// Number of queries in this database that have been canceled due to pinned buffers
	ConflBufferpin nullable.Int64 `json:"confl_bufferpin"`
	// Number of queries in this database that have been canceled due to deadlocks
	ConflDeadlock nullable.Int64 `json:"confl_deadlock"`
}

func (s *PgStats) fetchDatabaseConflicts() ([]PgStatDatabaseConflictsRow, error) {
	db := s.conn.db
	query := "select datid,datname," +
		"confl_tablespace,confl_lock,confl_snapshot,confl_bufferpin,confl_deadlock" +
		" from pg_stat_database_conflicts"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatDatabaseConflictsRow, 0)
	for rows.Next() {
		row := new(PgStatDatabaseConflictsRow)
		err := rows.Scan(&row.Datid, &row.Datname,
			&row.ConflTablespace, &row.ConflLock, &row.ConflSnapshot, &row.ConflBufferpin, &row.ConflDeadlock)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
