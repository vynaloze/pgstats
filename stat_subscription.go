package pgstats

import (
	"github.com/pkg/errors"
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatSubscriptionView reprowents content of pg_stat_subscription view
type PgStatSubscriptionView []PgStatSubscriptionRow

// PgStatSubscriptionRow reprowents schema of pg_stat_subscription view
type PgStatSubscriptionRow struct {
	// OID of the subscription
	Subid nullable.Int64 `json:"subid"`
	// Name of the subscription
	Subname nullable.String `json:"subname"`
	// Process ID of the subscription worker process
	Pid nullable.Int64 `json:"pid"`
	// OID of the relation that the worker is synchronizing; null for the main apply worker
	Relid nullable.Int64 `json:"relid"`
	// Last write-ahead log location received, the initial value of this field being 0
	ReceivedLsn nullable.Int64 `json:"received_lsn"`
	// Send time of last message received from origin WAL sender
	LastMsgSendTime nullable.Time `json:"last_msg_send_time"`
	// Receipt time of last message received from origin WAL sender
	LastMsgReceiptTime nullable.Time `json:"last_msg_receipt_time"`
	// Last write-ahead log location reported to origin WAL sender
	LatestEndLsn nullable.Int64 `json:"latest_end_lsn"`
	// Time of last write-ahead log location reported to origin WAL sender
	LatestEndTime nullable.Time `json:"latest_end_time"`
}

func (s *PgStats) fetchSubscription() (PgStatSubscriptionView, error) {
	version, err := s.getPgVersion()
	if err != nil {
		return nil, err
	}
	if version < 10 {
		return nil, errors.Errorf("Unsupported PostgreSQL version: %f", version)
	}

	db := s.conn.db
	query := "select subid,subname,pid,relid,received_lsn," +
		"last_msg_send_time,last_msg_receipt_time,latest_end_lsn,latest_end_time " +
		"from pg_stat_subscription"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatSubscriptionRow, 0)
	for rows.Next() {
		row := new(PgStatSubscriptionRow)
		err := rows.Scan(&row.Subid, &row.Subname, &row.Pid, &row.Relid, &row.ReceivedLsn,
			&row.LastMsgSendTime, &row.LastMsgReceiptTime, &row.LatestEndLsn, &row.LatestEndTime)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
