package pgstats

type PgStats struct {
	conn *connection
}

func Connect(dbname string, user string, password string, options ...func(*connection) error) (*PgStats, error) {
	s := &PgStats{}
	err := s.prepareConnection(dbname, user, password, options...)
	if err != nil {
		return nil, err
	}
	err = s.openConnection()
	return s, err
}

// PgStatAllIndexes returns an array containing statistics about accesses
// to each index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatAllIndexes() (PgStatAllIndexesView, error) {
	return s.fetchIndexes("pg_stat_all_indexes")
}

// PgStatUserIndexes returns an array containing statistics about accesses
// to each user-defined index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatUserIndexes() (PgStatUserIndexesView, error) {
	return s.fetchIndexes("pg_stat_user_indexes")
}

// PgStatSystemIndexes returns an array containing statistics about accesses
// to each system index in the current database.
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatSystemIndexes() (PgStatSystemIndexesView, error) {
	return s.fetchIndexes("pg_stat_sys_indexes")
}

// PgStatIoAllIndexes returns an array containing statistics about I/O
// on each index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoAllIndexes() (PgStatIoAllIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_all_indexes")
}

// PgStatIoUserIndexes returns an array containing statistics about I/O
// on each user-defined index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoUserIndexes() (PgStatIoUserIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_user_indexes")
}

// PgStatIoSystemIndexes returns an array containing statistics about I/O
// on each system index in the current database.
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoSystemIndexes() (PgStatIoSystemIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_sys_indexes")
}
