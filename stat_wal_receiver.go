package pgstats

import (
	"database/sql"
	"github.com/lib/pq"
)

// PgStatWalReceiverView represents content of pg_stat_wal_receiver view
type PgStatWalReceiverView struct {
	// Process ID of the WAL receiver process
	Pid int64 `json:"pid"`
	// Activity status of the WAL receiver process
	Status string `json:"status"`
	// First write-ahead log location used when WAL receiver is started
	ReceiveStartLsn sql.NullInt64 `json:"receive_start_lsn"`
	// First timeline number used when WAL receiver is started
	ReceiveStartTli sql.NullInt64 `json:"receive_start_tli"`
	// Last write-ahead log location already received and flushed to disk,
	// the initial value of this field being the first log location used when WAL receiver is started
	ReceivedLsn sql.NullInt64 `json:"received_lsn"`
	// Timeline number of last write-ahead log location received and flushed to disk,
	// the initial value of this field being the timeline number of the first log location used when WAL receiver is started
	ReceivedTli sql.NullInt64 `json:"received_tli"`
	// Send time of last message received from origin WAL sender
	LastMsgSendTime pq.NullTime `json:"last_msg_send_time"`
	// Receipt time of last message received from origin WAL sender
	LastMsgReceiptTime pq.NullTime `json:"last_msg_receipt_time"`
	// Last write-ahead log location reported to origin WAL sender
	LatestEndLsn sql.NullInt64 `json:"latest_end_lsn"`
	// Time of last write-ahead log location reported to origin WAL sender
	LatestEndTime pq.NullTime `json:"latest_end_time"`
	// Replication slot name used by this WAL receiver
	SlotName sql.NullString `json:"slot_name"`
	// Host of the PostgreSQL instance this WAL receiver is connected to.
	// This can be a host name, an IP address, or a directory path if the connection is via Unix socket.
	// (The path case can be distinguished because it will always be an absolute path, beginning with /.)
	SenderHost sql.NullString `json:"sender_host"`
	// Port number of the PostgreSQL instance this WAL receiver is connected to.
	SenderPort sql.NullInt64 `json:"sender_port"`
	// Connection string used by this WAL receiver, with security-sensitive fields obfuscated.
	Conninfo sql.NullString `json:"conninfo"`
}

func (s *PgStats) fetchWalReceiver() (PgStatWalReceiverView, error) {
	db := s.conn.db
	query := "select pid,status,receive_start_lsn,receive_start_tli,received_lsn," +
		"received_tli,last_msg_send_time,last_msg_receipt_time,latest_end_lsn,latest_end_time," +
		"slot_name,sender_host,sender_port,conninfo from pg_stat_wal_receiver"
	row := db.QueryRow(query)
	res := new(PgStatWalReceiverView)
	err := row.Scan(&res.Pid, &res.Status, &res.ReceiveStartLsn, &res.ReceiveStartTli, &res.ReceivedLsn,
		&res.ReceivedTli, &res.LastMsgSendTime, &res.LastMsgReceiptTime, &res.LatestEndLsn, &res.LatestEndTime,
		&res.SlotName, &res.SenderHost, &res.SenderPort, &res.Conninfo)
	return *res, err
}
