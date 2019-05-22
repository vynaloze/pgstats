package pgstats

import (
	"database/sql"
	"github.com/lib/pq"
)

// PgStatActivityView represents content of pg_stat_activity view
type PgStatActivityView []PgStatActivityRow

// PgStatActivityRow represents schema of pg_stat_activity view
type PgStatActivityRow struct {
	// OID of the database this backend is connected to
	Datid sql.NullInt64 `json:"datid"`
	// Name of the database this backend is connected to
	Datname sql.NullString `json:"datname"`
	// Process ID of this backend
	Pid int64 `json:"pid"`
	// OID of the user logged into this backend
	Usesysid sql.NullInt64 `json:"usesysid"`
	// Name of the user logged into this backend
	Usename sql.NullString `json:"usename"`
	// Name of the application that is connected to this backend
	ApplicationName sql.NullString `json:"application_name"`
	// IP address of the client connected to this backend.
	// If this field is null, it indicates either that the client is connected via a Unix socket on the server machine
	// or that this is an internal process such as autovacuum.
	ClientAddr sql.NullString `json:"client_addr"`
	// Host name of the connected client, as reported by a reverse DNS lookup of client_addr.
	// This field will only be non-null for IP connections, and only when log_hostname is enabled.
	ClientHostname sql.NullString `json:"client_hostname"`
	// TCP port number that the client is using for communication with this backend,
	// or -1 if a Unix socket is used
	ClientPort sql.NullInt64 `json:"client_port"`
	// Time when this process was started.
	// For client backends, this is the time the client connected to the server.
	BackendStart pq.NullTime `json:"backend_start"`
	// Time when this process' current transaction was started,
	// or null if no transaction is active.
	// If the current query is the first of its transaction, this column is equal to the query_start column.
	XactStart pq.NullTime `json:"xact_start"`
	// Time when the currently active query was started,
	// or if state is not active, when the last query was started
	QueryStart pq.NullTime `json:"query_start"`
	// ime when the state was last changed
	StateChange pq.NullTime `json:"state_change"`
	// The type of event for which the backend is waiting, if any; otherwise NULL.
	// Supported since PostgreSQL 9.6.
	// For possible values, see:
	// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ACTIVITY-VIEW
	WaitEventType sql.NullString `json:"wait_event_type"`
	// Wait event name if backend is currently waiting, otherwise NULL.
	// Supported since PostgreSQL 9.6.
	// For details, see:
	// https://www.postgresql.org/docs/current/monitoring-stats.html#WAIT-EVENT-TABLE
	WaitEvent sql.NullString `json:"wait_event"`
	// True if this backend is currently waiting on a lock.
	// Supported until PostgreSQL 9.5 (inclusive).
	Waiting sql.NullBool `json:"waiting"`
	// Current overall state of this backend.
	// For possible values, see:
	// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-ACTIVITY-VIEW
	State sql.NullString `json:"state"`
	// Top-level transaction identifier of this backend, if any.
	BackendXid sql.NullInt64 `json:"backend_xid"`
	// The current backend's xmin horizon.
	BackendXmin sql.NullInt64 `json:"backend_xmin"`
	// Text of this backend's most recent query.
	// If state is active this field shows the currently executing query.
	// In all other states, it shows the last query that was executed.
	// By default the query text is truncated at 1024 characters;
	// this value can be changed via the parameter track_activity_query_size.
	Query sql.NullString `json:"query"`
	// Type of current backend.
	// Possible types are autovacuum launcher, autovacuum worker, logical replication launcher,
	// logical replication worker, parallel worker, background writer, client backend, checkpointer,
	// startup, walreceiver, walsender and walwriter.
	// In addition, background workers registered by extensions may have additional types.
	// Supported since PostgreSQL 10
	BackendType sql.NullString `json:"backend_type"`
}

func (s *PgStats) fetchActivity() ([]PgStatActivityRow, error) {
	version, err := s.getPgVersion()
	if err != nil {
		return nil, err
	}
	if version > 9.6 {
		return s.fetchActivity10()
	}
	if version == 9.6 {
		return s.fetchActivity96()
	}
	return s.fetchActivity95()
}

func (s *PgStats) fetchActivity10() ([]PgStatActivityRow, error) {
	db := s.conn.db
	query := "select datid,datname,pid,usesysid,usename," +
		"application_name,client_addr,client_hostname,client_port,backend_start," +
		"xact_start,query_start,state_change,wait_event_type,wait_event," +
		"state,backend_xid,backend_xmin,query,backend_type from pg_stat_activity"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatActivityRow, 0)
	for rows.Next() {
		row := new(PgStatActivityRow)
		err := rows.Scan(&row.Datid, &row.Datname, &row.Pid, &row.Usesysid, &row.Usename,
			&row.ApplicationName, &row.ClientAddr, &row.ClientHostname, &row.ClientPort, &row.BackendStart,
			&row.XactStart, &row.QueryStart, &row.StateChange, &row.WaitEventType, &row.WaitEvent,
			&row.State, &row.BackendXid, &row.BackendXmin, &row.Query, &row.BackendType)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}

func (s *PgStats) fetchActivity96() ([]PgStatActivityRow, error) {
	db := s.conn.db
	query := "select datid,datname,pid,usesysid,usename," +
		"application_name,client_addr,client_hostname,client_port,backend_start," +
		"xact_start,query_start,state_change,wait_event_type,wait_event," +
		"state,backend_xid,backend_xmin,query from pg_stat_activity"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatActivityRow, 0)
	for rows.Next() {
		row := new(PgStatActivityRow)
		err := rows.Scan(&row.Datid, &row.Datname, &row.Pid, &row.Usesysid, &row.Usename,
			&row.ApplicationName, &row.ClientAddr, &row.ClientHostname, &row.ClientPort, &row.BackendStart,
			&row.XactStart, &row.QueryStart, &row.StateChange, &row.WaitEventType, &row.WaitEvent,
			&row.State, &row.BackendXid, &row.BackendXmin, &row.Query)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}

func (s *PgStats) fetchActivity95() ([]PgStatActivityRow, error) {
	db := s.conn.db
	query := "select datid,datname,pid,usesysid,usename," +
		"application_name,client_addr,client_hostname,client_port,backend_start," +
		"xact_start,query_start,state_change,waiting," +
		"state,backend_xid,backend_xmin,query from pg_stat_activity"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatActivityRow, 0)
	for rows.Next() {
		row := new(PgStatActivityRow)
		err := rows.Scan(&row.Datid, &row.Datname, &row.Pid, &row.Usesysid, &row.Usename,
			&row.ApplicationName, &row.ClientAddr, &row.ClientHostname, &row.ClientPort, &row.BackendStart,
			&row.XactStart, &row.QueryStart, &row.StateChange, &row.Waiting,
			&row.State, &row.BackendXid, &row.BackendXmin, &row.Query)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
