// Package pgstats provides convenient access to all pg_stat_* statistics,
// allowing to monitor PostgreSQL instances inside go applications.
//
// For details, see: https://github.com/vynaloze/pgstats/blob/master/README.md
package pgstats

// PgStats holds a single connection to the database
// and provides a convenient access to all postgres monitoring statistics.
type PgStats struct {
	conn *connection
}

// Connect opens a connection using provided parameters and returns a pointer to newly created PgStats struct.
func Connect(dbname string, user string, password string, options ...Option) (*PgStats, error) {
	s := &PgStats{}
	err := s.prepareConnection(dbname, user, password, options...)
	if err != nil {
		return nil, err
	}
	err = s.openConnection()
	return s, err
}

// Close closes the connection to database.
func (s *PgStats) Close() error {
	return s.conn.db.Close()
}

// PgStatActivity returns a slice, containing information related to the current activity of a process,
// such as state and current query, for each server process.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ACTIVITY-VIEW
func (s *PgStats) PgStatActivity() (PgStatActivityView, error) {
	return s.fetchActivity()
}

// PgStatReplication returns a slice, containing statistics about each WAL sender process,
// showing information about replication to that sender's connected standby server.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-REPLICATION-VIEW
func (s *PgStats) PgStatReplication() (PgStatReplicationView, error) {
	return s.fetchReplication()
}

// PgStatWalReceiver returns a single struct,
// containing statistics about the WAL receiver from that receiver's connected server.
//
// Supported since PostgreSQL 9.6.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-WAL-RECEIVER-VIEW
func (s *PgStats) PgStatWalReceiver() (PgStatWalReceiverView, error) {
	return s.fetchWalReceiver()
}

// PgStatSubscription returns a slice, containing statistics about
// subscription for main worker (with null PID if the worker is not running),
// and workers handling the initial data copy of the subscribed tables.
//
// Supported since PostgreSQL 10.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-SUBSCRIPTION
func (s *PgStats) PgStatSubscription() (PgStatSubscriptionView, error) {
	return s.fetchSubscription()
}

// PgStatSsl returns a slice, containing statistics about SSL usage
// on the connection for each backend or WAL sender process.
//
// Supported since PostgreSQL 9.5.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-SSL
func (s *PgStats) PgStatSsl() (PgStatSslView, error) {
	return s.fetchSsl()
}

// PgStatProgressVacuum returns a slice, containing information related to currently running VACUUM processes,
// for each backend (including autovacuum worker processes) that is currently vacuuming.
// Progress reporting is not currently supported for VACUUM FULL and backends running VACUUM FULL will not be listed in this view.
//
// Supported since PostgreSQL 9.6.
//
// For more details, see:
// https://www.postgresql.org/docs/current/progress-reporting.html#VACUUM-PROGRESS-REPORTING
func (s *PgStats) PgStatProgressVacuum() (PgStatProgressVacuumView, error) {
	return s.fetchProgressVacuum()
}

// PgStatArchiver returns a single struct, containing global data for the cluster,
// showing statistics about the WAL archiver process's activity.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ARCHIVER-VIEW
func (s *PgStats) PgStatArchiver() (PgStatArchiverView, error) {
	return s.fetchArchiver()
}

// PgStatBgWriter returns a single struct, containing global data for the cluster,
// showing statistics about the background writer process's activity.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-BGWRITER-VIEW
func (s *PgStats) PgStatBgWriter() (PgStatBgWriterView, error) {
	return s.fetchBgWriter()
}

// PgStatDatabase returns a slice containing database-wide statistics for each database in the cluster.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-DATABASE-VIEW
func (s *PgStats) PgStatDatabase() (PgStatDatabaseView, error) {
	return s.fetchDatabases()
}

// PgStatDatabaseConflicts returns a slice containing database-wide statistics for each database in the cluster about
// query cancels occurring due to conflicts with recovery on standby servers.
// This will only contain information on standby servers, since conflicts do not occur on master servers.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-DATABASE-CONFLICTS-VIEW
func (s *PgStats) PgStatDatabaseConflicts() (PgStatDatabaseConflictsView, error) {
	return s.fetchDatabaseConflicts()
}

// PgStatAllTables returns a slice containing statistics about accesses
// to each table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func (s *PgStats) PgStatAllTables() (PgStatAllTablesView, error) {
	return s.fetchTables("pg_stat_all_tables")
}

// PgStatSystemTables returns a slice containing statistics about accesses
// to each system table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func (s *PgStats) PgStatSystemTables() (PgStatSystemTablesView, error) {
	return s.fetchTables("pg_stat_sys_tables")
}

// PgStatUserTables returns a slice containing statistics about accesses
// to each user-defined table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func (s *PgStats) PgStatUserTables() (PgStatUserTablesView, error) {
	return s.fetchTables("pg_stat_user_tables")
}

// PgStatXactAllTables returns a slice containing statistics about accesses
// to each table in the current database (including TOAST tables),
// but counts only actions taken so far within the current transaction
// (which are not yet included in pg_stat_all_tables and related views).
func (s *PgStats) PgStatXactAllTables() (PgStatXactAllTablesView, error) {
	return s.fetchXactTables("pg_stat_xact_all_tables")
}

// PgStatXactSystemTables returns a slice containing statistics about accesses
// to each system table in the current database (including TOAST tables),
// but counts only actions taken so far within the current transaction
// (which are not yet included in pg_stat_all_tables and related views).
func (s *PgStats) PgStatXactSystemTables() (PgStatXactSystemTablesView, error) {
	return s.fetchXactTables("pg_stat_xact_sys_tables")
}

// PgStatXactUserTables returns a slice containing statistics about accesses
// to each user-defined table in the current database (including TOAST tables),
// but counts only actions taken so far within the current transaction
// (which are not yet included in pg_stat_all_tables and related views).
func (s *PgStats) PgStatXactUserTables() (PgStatXactUserTablesView, error) {
	return s.fetchXactTables("pg_stat_xact_user_tables")
}

// PgStatAllIndexes returns a slice containing statistics about accesses
// to each index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatAllIndexes() (PgStatAllIndexesView, error) {
	return s.fetchIndexes("pg_stat_all_indexes")
}

// PgStatSystemIndexes returns a slice containing statistics about accesses
// to each system index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatSystemIndexes() (PgStatSystemIndexesView, error) {
	return s.fetchIndexes("pg_stat_sys_indexes")
}

// PgStatUserIndexes returns a slice containing statistics about accesses
// to each user-defined index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func (s *PgStats) PgStatUserIndexes() (PgStatUserIndexesView, error) {
	return s.fetchIndexes("pg_stat_user_indexes")
}

// PgStatIoAllTables returns a slice containing statistics about I/O
// on each table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func (s *PgStats) PgStatIoAllTables() (PgStatIoAllTablesView, error) {
	return s.fetchIoTables("pg_statio_all_tables")
}

// PgStatIoSystemTables returns a slice containing statistics about I/O
// on each system table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func (s *PgStats) PgStatIoSystemTables() (PgStatIoSystemTablesView, error) {
	return s.fetchIoTables("pg_statio_sys_tables")
}

// PgStatIoUserTables returns a slice containing statistics about I/O
// on each user-defined table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func (s *PgStats) PgStatIoUserTables() (PgStatIoUserTablesView, error) {
	return s.fetchIoTables("pg_statio_user_tables")
}

// PgStatIoAllIndexes returns a slice containing statistics about I/O
// on each index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoAllIndexes() (PgStatIoAllIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_all_indexes")
}

// PgStatIoSystemIndexes returns a slice containing statistics about I/O
// on each system index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoSystemIndexes() (PgStatIoSystemIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_sys_indexes")
}

// PgStatIoUserIndexes returns a slice containing statistics about I/O
// on each user-defined index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func (s *PgStats) PgStatIoUserIndexes() (PgStatIoUserIndexesView, error) {
	return s.fetchIoIndexes("pg_statio_user_indexes")
}

// PgStatIoAllSequences returns a slice containing statistics about I/O
// on each sequence in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-SEQUENCES-VIEW
func (s *PgStats) PgStatIoAllSequences() (PgStatIoAllSequencesView, error) {
	return s.fetchIoSequences("pg_statio_all_sequences")
}

// PgStatIoSystemSequences returns a slice containing statistics about I/O
// on each system sequence in the current database.
// (As of PostgreSQL 11, no system sequences are defined, so this view is always empty.)
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-SEQUENCES-VIEW
func (s *PgStats) PgStatIoSystemSequences() (PgStatIoSystemSequencesView, error) {
	return s.fetchIoSequences("pg_statio_sys_sequences")
}

// PgStatIoUserSequences returns a slice containing statistics about I/O
// on each user-defined sequence in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-SEQUENCES-VIEW
func (s *PgStats) PgStatIoUserSequences() (PgStatIoUserSequencesView, error) {
	return s.fetchIoSequences("pg_statio_user_sequences")
}

// PgStatUserFunctions returns a slice containing statistics about executions
// of each tracked function in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-USER-FUNCTIONS-VIEW
func (s *PgStats) PgStatUserFunctions() (PgStatUserFunctionsView, error) {
	return s.fetchFunctions("pg_stat_user_functions")
}

// PgStatXactUserFunctions returns a slice containing statistics about executions
// of each tracked function in the current database,
// but counts only calls during the current transaction
// (which are not yet included in pg_stat_user_functions).
func (s *PgStats) PgStatXactUserFunctions() (PgStatXactUserFunctionsView, error) {
	return s.fetchFunctions("pg_stat_xact_user_functions")
}

// PgStatStatements returns a slice containing statistics about executions
// of all SQL statements in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/pgstatstatements.html
func (s *PgStats) PgStatStatements() (PgStatStatementsView, error) {
	return s.fetchStatements()
}
