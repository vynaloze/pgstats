package pgstats

import "database/sql"

// PgStatProgressVacuumView represents content of pg_stat_progress_vacuum view
type PgStatProgressVacuumView []PgStatProgressVacuumRow

// PgStatProgressVacuumRow represents schema of pg_stat_progress_vacuum view
type PgStatProgressVacuumRow struct {
	// Process ID of backend.
	Pid int64 `json:"pid"`
	// OID of the database to which this backend is connected.
	Datid int64 `json:"datid"`
	// Name of the database to which this backend is connected.
	Datname string `json:"datname"`
	// OID of the table being vacuumed.
	Relid int64 `json:"relid"`
	// Current processing phase of vacuum. See:
	// https://www.postgresql.org/docs/current/progress-reporting.html#VACUUM-PHASES
	Phase string `json:"phase"`
	// Total number of heap blocks in the table. This number is reported as of the beginning of the scan;
	// blocks added later will not be (and need not be) visited by this VACUUM.
	HeapBlksTotal sql.NullInt64 `json:"heap_blks_total"`
	// Number of heap blocks scanned.
	// Because the visibility map is used to optimize scans, some blocks will be skipped without inspection;
	// skipped blocks are included in this total, so that this number will eventually become equal to heap_blks_total
	// when the vacuum is complete. This counter only advances when the phase is scanning heap.
	HeapBlksScanned sql.NullInt64 `json:"heap_blks_scanned"`
	// Number of heap blocks vacuumed. Unless the table has no indexes, this counter only advances when the phase is vacuuming heap.
	// Blocks that contain no dead tuples are skipped, so the counter may sometimes skip forward in large increments.
	HeapBlksVacuumed sql.NullInt64 `json:"heap_blks_vacuumed"`
	// Number of completed index vacuum cycles.
	IndexVacuumCount sql.NullInt64 `json:"index_vacuum_count"`
	// Number of dead tuples that we can store before needing to perform an index vacuum cycle, based on maintenance_work_mem.
	MaxDeadTuples sql.NullInt64 `json:"max_dead_tuples"`
	// Number of dead tuples collected since the last index vacuum cycle.
	NumDeadTuples sql.NullInt64 `json:"num_dead_tuples"`
}

func (s *PgStats) fetchProgressVacuum() (PgStatProgressVacuumView, error) {
	db := s.conn.db
	query := "select pid,datid,datname,relid,phase," +
		"heap_blks_total,heap_blks_scanned,heap_blks_vacuumed,index_vacuum_count,max_dead_tuples," +
		"num_dead_tuples from pg_stat_progress_vacuum"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(PgStatProgressVacuumView, 0)
	for rows.Next() {
		row := new(PgStatProgressVacuumRow)
		err := rows.Scan(&row.Pid, &row.Datid, &row.Datname, &row.Relid, &row.Phase,
			&row.HeapBlksTotal, &row.HeapBlksScanned, &row.HeapBlksVacuumed, &row.IndexVacuumCount, &row.MaxDeadTuples,
			&row.NumDeadTuples)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
