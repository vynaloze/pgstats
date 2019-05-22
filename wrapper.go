package pgstats

import (
	"errors"
	"sync"
)

var wrapper = struct {
	stats  *PgStats
	once   sync.Once
	opened bool
}{
	&PgStats{},
	sync.Once{},
	false,
}

// DefineConnection defines the connection, which can be later used globally to collect statistics.
func DefineConnection(dbname string, user string, password string, options ...func(*connection) error) error {
	var err error
	wrapper.once.Do(func() {
		wrapper.stats, err = Connect(dbname, user, password, options...)
		wrapper.opened = true
	})
	return err
}

// PgStatActivity returns a slice, containing information related to the current activity of a process,
// such as state and current query, for each server process.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ACTIVITY-VIEW
func PgStatActivity() (PgStatActivityView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchActivity()
}

// PgStatReplication returns a slice, containing statistics about each WAL sender process,
// showing information about replication to that sender's connected standby server.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-REPLICATION-VIEW
func PgStatReplication() (PgStatReplicationView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchReplication()
}

// PgStatWalReceiver returns a single struct,
// containing statistics about the WAL receiver from that receiver's connected server.
//
// Supported since PostgreSQL 9.6.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-WAL-RECEIVER-VIEW
func PgStatWalReceiver() (PgStatWalReceiverView, error) {
	if !wrapper.opened {
		return PgStatWalReceiverView{}, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchWalReceiver()
}

// PgStatSubscription returns a slice, containing statistics about
// subscription for main worker (with null PID if the worker is not running),
// and workers handling the initial data copy of the subscribed tables.
//
// Supported since PostgreSQL 10.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-SUBSCRIPTION
func PgStatSubscription() (PgStatSubscriptionView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchSubscription()
}

// PgStatSsl returns a slice, containing statistics about SSL usage
// on the connection for each backend or WAL sender process.
//
// Supported since PostgreSQL 9.5.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-SSL
func PgStatSsl() (PgStatSslView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchSsl()
}

// PgStatProgressVacuum returns a slice, containing information related to currently running VACUUM processes,
// for each backend (including autovacuum worker processes) that is currently vacuuming.
// Progress reporting is not currently supported for VACUUM FULL and backends running VACUUM FULL will not be listed in this view.
//
// Supported since PostgreSQL 9.6.
//
// For more details, see:
// https://www.postgresql.org/docs/current/progress-reporting.html#VACUUM-PROGRESS-REPORTING
func PgStatProgressVacuum() (PgStatProgressVacuumView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchProgressVacuum()
}

// PgStatArchiver returns a single struct, containing global data for the cluster,
// showing statistics about the WAL archiver process's activity.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ARCHIVER-VIEW
func PgStatArchiver() (PgStatArchiverView, error) {
	if !wrapper.opened {
		return PgStatArchiverView{}, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchArchiver()
}

// PgStatBgWriter returns a single struct, containing global data for the cluster,
// showing statistics about the background writer process's activity.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-BGWRITER-VIEW
func PgStatBgWriter() (PgStatBgWriterView, error) {
	if !wrapper.opened {
		return PgStatBgWriterView{}, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchBgWriter()
}

// PgStatDatabase returns a slice containing database-wide statistics for each database in the cluster.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-DATABASE-VIEW
func PgStatDatabase() (PgStatDatabaseView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchDatabases()
}

// PgStatDatabaseConflicts returns a slice containing database-wide statistics for each database in the cluster about
// query cancels occurring due to conflicts with recovery on standby servers.
// This will only contain information on standby servers, since conflicts do not occur on master servers.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-DATABASE-CONFLICTS-VIEW
func PgStatDatabaseConflicts() (PgStatDatabaseConflictsView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchDatabaseConflicts()
}

// PgStatAllTables returns a slice containing statistics about accesses
// to each table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func PgStatAllTables() (PgStatAllTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchTables("pg_stat_all_tables")
}

// PgStatSystemTables returns a slice containing statistics about accesses
// to each system table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func PgStatSystemTables() (PgStatSystemTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchTables("pg_stat_sys_tables")
}

// PgStatUserTables returns a slice containing statistics about accesses
// to each user-defined table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-TABLES-VIEW
func PgStatUserTables() (PgStatUserTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchTables("pg_stat_user_tables")
}

// PgStatXactAllTables returns a slice containing statistics about accesses
// to each table in the current database (including TOAST tables),
// but counts only actions taken so far within the current transaction
// (which are not yet included in pg_stat_all_tables and related views).
func PgStatXactAllTables() (PgStatXactAllTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchXactTables("pg_stat_xact_all_tables")
}

// PgStatXactSystemTables returns a slice containing statistics about accesses
// to each system table in the current database (including TOAST tables),
// but counts only actions taken so far within the current transaction
// (which are not yet included in pg_stat_all_tables and related views).
func PgStatXactSystemTables() (PgStatXactSystemTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchXactTables("pg_stat_xact_sys_tables")
}

// PgStatXactUserTables returns a slice containing statistics about accesses
// to each user-defined table in the current database (including TOAST tables),
// but counts only actions taken so far within the current transaction
// (which are not yet included in pg_stat_all_tables and related views).
func PgStatXactUserTables() (PgStatXactUserTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchXactTables("pg_stat_xact_user_tables")
}

// PgStatAllIndexes returns a slice containing statistics about accesses
// to each index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func PgStatAllIndexes() (PgStatAllIndexesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIndexes("pg_stat_all_indexes")
}

// PgStatSystemIndexes returns a slice containing statistics about accesses
// to each system index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func PgStatSystemIndexes() (PgStatSystemIndexesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIndexes("pg_stat_sys_indexes")
}

// PgStatUserIndexes returns a slice containing statistics about accesses
// to each user-defined index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ALL-INDEXES-VIEW
func PgStatUserIndexes() (PgStatUserIndexesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIndexes("pg_stat_user_indexes")
}

// PgStatIoAllTables returns a slice containing statistics about I/O
// on each table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func PgStatIoAllTables() (PgStatIoAllTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoTables("pg_statio_all_tables")
}

// PgStatIoSystemTables returns a slice containing statistics about I/O
// on each system table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func PgStatIoSystemTables() (PgStatIoSystemTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoTables("pg_statio_sys_tables")
}

// PgStatIoUserTables returns a slice containing statistics about I/O
// on each user-defined table in the current database (including TOAST tables).
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-TABLES-VIEW
func PgStatIoUserTables() (PgStatIoUserTablesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoTables("pg_statio_user_tables")
}

// PgStatIoAllIndexes returns a slice containing statistics about I/O
// on each index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func PgStatIoAllIndexes() (PgStatIoAllIndexesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoIndexes("pg_statio_all_indexes")
}

// PgStatIoSystemIndexes returns a slice containing statistics about I/O
// on each system index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func PgStatIoSystemIndexes() (PgStatIoSystemIndexesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoIndexes("pg_statio_sys_indexes")
}

// PgStatIoUserIndexes returns a slice containing statistics about I/O
// on each user-defined index in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-INDEXES-VIEW
func PgStatIoUserIndexes() (PgStatIoUserIndexesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoIndexes("pg_statio_user_indexes")
}

// PgStatIoAllSequences returns a slice containing statistics about I/O
// on each sequence in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-SEQUENCES-VIEW
func PgStatIoAllSequences() (PgStatIoAllSequencesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoSequences("pg_statio_all_sequences")
}

// PgStatIoSystemSequences returns a slice containing statistics about I/O
// on each system sequence in the current database.
// (As of PostgreSQL 11, no system sequences are defined, so this view is always empty.)
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-SEQUENCES-VIEW
func PgStatIoSystemSequences() (PgStatIoSystemSequencesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoSequences("pg_statio_sys_sequences")
}

// PgStatIoUserSequences returns a slice containing statistics about I/O
// on each user-defined sequence in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STATIO-ALL-SEQUENCES-VIEW
func PgStatIoUserSequences() (PgStatIoUserSequencesView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchIoSequences("pg_statio_user_sequences")
}

// PgStatUserFunctions returns a slice containing statistics about executions
// of each tracked function in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-USER-FUNCTIONS-VIEW
func PgStatUserFunctions() (PgStatUserFunctionsView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchFunctions("pg_stat_user_functions")
}

// PgStatXactUserFunctions returns a slice containing statistics about executions
// of each tracked function in the current database,
// but counts only calls during the current transaction
// (which are not yet included in pg_stat_user_functions).
func PgStatXactUserFunctions() (PgStatXactUserFunctionsView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchFunctions("pg_stat_xact_user_functions")
}

// PgStatStatements returns a slice containing statistics about executions
// of all SQL statements in the current database.
//
// For more details, see:
// https://www.postgresql.org/docs/current/pgstatstatements.html
func PgStatStatements() (PgStatStatementsView, error) {
	if !wrapper.opened {
		return nil, errors.New("connection has not been defined")
	}
	return wrapper.stats.fetchStatements()
}
