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

// PgStatAllTables returns an array containing statistics about accesses
// to each table in the current database (including TOAST tables).
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func (s *PgStats) PgStatAllTables() (PgStatAllTablesView, error) {
	return s.fetchTables("pg_stat_all_tables")
}

// PgStatSystemTables returns an array containing statistics about accesses
// to each system table in the current database (including TOAST tables).
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func (s *PgStats) PgStatSystemTables() (PgStatSystemTablesView, error) {
	return s.fetchTables("pg_stat_sys_tables")
}

// PgStatUserTables returns an array containing statistics about accesses
// to each user-defined table in the current database (including TOAST tables).
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func (s *PgStats) PgStatUserTables() (PgStatUserTablesView, error) {
	return s.fetchTables("pg_stat_user_tables")
}

// PgStatAllIndexes returns an array containing statistics about accesses
// to each index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatAllIndexes() (PgStatAllIndexesView, error) {
	return s.fetchIndexes("pg_stat_all_indexes")
}

// PgStatSystemIndexes returns an array containing statistics about accesses
// to each system index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatSystemIndexes() (PgStatSystemIndexesView, error) {
	return s.fetchIndexes("pg_stat_sys_indexes")
}

// PgStatUserIndexes returns an array containing statistics about accesses
// to each user-defined index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatUserIndexes() (PgStatUserIndexesView, error) {
	return s.fetchIndexes("pg_stat_user_indexes")
}

// PgStatIoAllTables returns an array containing statistics about I/O
// on each table in the current database (including TOAST tables).
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func (s *PgStats) PgStatIoAllTables() (PgStatIoAllTablesView, error) {
	return s.fetchIoTables("pg_statio_all_tables")
}

// PgStatIoSystemTables returns an array containing statistics about I/O
// on each system table in the current database (including TOAST tables).
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func (s *PgStats) PgStatIoSystemTables() (PgStatIoSystemTablesView, error) {
	return s.fetchIoTables("pg_statio_sys_tables")
}

// PgStatIoUserTables returns an array containing statistics about I/O
// on each user-defined table in the current database (including TOAST tables).
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func (s *PgStats) PgStatIoUserTables() (PgStatIoUserTablesView, error) {
	return s.fetchIoTables("pg_statio_user_tables")
}

// PgStatIoAllIndexes returns an array containing statistics about I/O
// on each index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoAllIndexes() (PgStatIoAllIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_all_indexes")
}

// PgStatIoSystemIndexes returns an array containing statistics about I/O
// on each system index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoSystemIndexes() (PgStatIoSystemIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_sys_indexes")
}

// PgStatIoUserIndexes returns an array containing statistics about I/O
// on each user-defined index in the current database.
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoUserIndexes() (PgStatIoUserIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_user_indexes")
}
