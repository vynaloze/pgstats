package pgstats

import (
	"database/sql"
	"github.com/lib/pq"
)

// PgStatReplicationView represents content of pg_stat_replication view
type PgStatReplicationView []PgStatReplicationRow

// PgStatReplicationRow represents schema of pg_stat_replication view
type PgStatReplicationRow struct {
	// Process ID of a WAL sender process
	Pid int64 `json:"pid"`
	// OID of the user logged into this WAL sender process
	Usesysid sql.NullInt64 `json:"usesysid"`
	// Name of the user logged into this WAL sender process
	Usename sql.NullString `json:"usename"`
	// Name of the application that is connected to this WAL sender
	ApplicationName sql.NullString `json:"application_name"`
	// IP address of the client connected to this WAL sender.
	// If this field is null, it indicates that the client is connected via a Unix socket on the server machine.
	ClientAddr sql.NullString `json:"client_addr"`
	// Host name of the connected client, as reported by a reverse DNS lookup of client_addr.
	// This field will only be non-null for IP connections, and only when log_hostname is enabled.
	ClientHostname sql.NullString `json:"client_hostname"`
	// TCP port number that the client is using for communication with this WAL sender, or -1 if a Unix socket is used
	ClientPort sql.NullInt64 `json:"client_port"`
	// Time when this process was started, i.e., when the client connected to this WAL sender
	BackendStart pq.NullTime `json:"backend_start"`
	// This standby's xmin horizon reported by hot_standby_feedback - see:
	// https://www.postgresql.org/docs/current/runtime-config-replication.html#GUC-HOT-STANDBY-FEEDBACK
	BackendXmin sql.NullInt64 `json:"backend_xmin"`
	// Current WAL sender state.
	// For possible values, see:
	// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-REPLICATION-VIEW
	State sql.NullString `json:"state"`
	// Last write-ahead log location sent on this connection
	SentLsn sql.NullInt64 `json:"sent_lsn"`
	// Last write-ahead log location written to disk by this standby server
	WriteLsn sql.NullInt64 `json:"write_lsn"`
	// Last write-ahead log location flushed to disk by this standby server
	FlushLsn sql.NullInt64 `json:"flush_lsn"`
	// Last write-ahead log location replayed into the database on this standby server
	ReplayLsn sql.NullInt64 `json:"replay_lsn"`
	// Time elapsed between flushing recent WAL locally and receiving notification that this standby server
	// has written it (but not yet flushed it or applied it). This can be used to gauge the delay
	// that synchronous_commit level remote_write incurred while committing
	// if this server was configured as a synchronous standby.
	WriteLag pq.NullTime `json:"write_lag"` // fixme?
	// Time elapsed between flushing recent WAL locally and receiving notification that this standby server
	// has written 	// and flushed it (but not yet applied it). This can be used to gauge the delay
	// that synchronous_commit level on incurred while committing
	// if this server was configured as a synchronous standby.
	FlushLag pq.NullTime `json:"flush_lag"` // fixme?
	// Time elapsed between flushing recent WAL locally and receiving notification that this standby server
	// has written, flushed and applied it. This can be used to gauge the delay
	// that synchronous_commit level remote_apply incurred while committing
	// if this server was configured as a synchronous standby.
	ReplayLag pq.NullTime `json:"replay_lag"` // fixme?
	// Priority of this standby server for being chosen as the synchronous standby
	// in a priority-based synchronous replication. This has no effect in a quorum-based synchronous replication.
	SyncPriority sql.NullInt64 `json:"sync_priority"`
	// Synchronous state of this standby server.
	// For possible values, see:
	// https://www.postgresql.org/docs/current/monitoring-stats.html#PG-STAT-REPLICATION-VIEW
	SyncState sql.NullString `json:"sync_state"`
}

func (s *PgStats) fetchReplication() ([]PgStatReplicationRow, error) {
	db := s.conn.db
	query := "select pid,usesysid,usename,application_name,client_addr," +
		"client_hostname,client_port,backend_start,backend_xmin,state," +
		"sent_lsn,write_lsn,flush_lsn,replay_lsn,write_lag," +
		"flush_lag,replay_lag,sync_priority,sync_state from pg_stat_replication"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatReplicationRow, 0)
	for rows.Next() {
		row := new(PgStatReplicationRow)
		err := rows.Scan(&row.Pid, &row.Usesysid, &row.Usename, &row.ApplicationName, &row.ClientAddr,
			&row.ClientHostname, &row.ClientPort, &row.BackendStart, &row.BackendXmin, &row.State,
			&row.SentLsn, &row.WriteLsn, &row.FlushLsn, &row.ReplayLsn, &row.WriteLag,
			&row.FlushLag, &row.ReplayLag, &row.SyncPriority, &row.SyncState)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
