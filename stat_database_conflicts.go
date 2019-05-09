package pgstats

import (
	"database/sql"
)

// PgStatDatabaseConflictsView represents content of pg_stat_database_conflicts view
type PgStatDatabaseConflictsView []PgStatDatabaseConflictsRow

// PgStatDatabaseConflictsRow represents schema of pg_stat_database_conflicts view
type PgStatDatabaseConflictsRow struct {
	// OID of a database
	DatId int64
	// Name of this database
	DatName string
	// Number of queries in this database that have been canceled due to dropped tablespaces
	ConflTablespace sql.NullInt64
	// Number of queries in this database that have been canceled due to lock timeouts
	ConflLock sql.NullInt64
	// Number of queries in this database that have been canceled due to old snapshots
	ConflSnapshot sql.NullInt64
	// Number of queries in this database that have been canceled due to pinned buffers
	ConflBufferpin sql.NullInt64
	// Number of queries in this database that have been canceled due to deadlocks
	ConflDeadlock sql.NullInt64
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
		err := rows.Scan(&row.DatId, &row.DatName,
			&row.ConflTablespace, &row.ConflLock, &row.ConflSnapshot, &row.ConflBufferpin, &row.ConflDeadlock)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
