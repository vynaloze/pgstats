package pgstats

import (
	"database/sql"
	"github.com/pkg/errors"
)

// PgStatSslView represents content of pg_stat_ssl view
type PgStatSslView []PgStatSslRow

// PgStatSslRow represents schema of pg_stat_ssl view
type PgStatSslRow struct {
	// Process ID of a backend or WAL sender process
	Pid int64 `json:"pid"`
	// True if SSL is used on this connection
	Ssl bool `json:"ssl"`
	// Version of SSL in use, or NULL if SSL is not in use on this connection
	Version sql.NullString `json:"version"`
	// Name of SSL cipher in use, or NULL if SSL is not in use on this connection
	Cipher sql.NullString `json:"cipher"`
	// Number of bits in the encryption algorithm used, or NULL if SSL is not used on this connection
	Bits sql.NullInt64 `json:"bits"`
	// True if SSL compression is in use, false if not, or NULL if SSL is not in use on this connection
	Compression sql.NullBool `json:"compression"`
	// Distinguished Name (DN) field from the client certificate used,
	// or NULL if no client certificate was supplied or if SSL is not in use on this connection.
	// This field is truncated if the DN field is longer than NAMEDATALEN (64 characters in a standard build)
	Clientdn sql.NullString `json:"clientdn"`
}

func (s *PgStats) fetchSsl() (PgStatSslView, error) {
	version, err := s.getPgVersion()
	if err != nil {
		return nil, err
	}
	if version < 9.5 {
		return nil, errors.Errorf("Unsupported PostgreSQL version: %f", version)
	}

	db := s.conn.db
	query := "select pid,ssl,version,cipher,bits," +
		"compression,clientdn from pg_stat_ssl"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(PgStatSslView, 0)
	for rows.Next() {
		row := new(PgStatSslRow)
		err := rows.Scan(&row.Pid, &row.Ssl, &row.Version, &row.Cipher, &row.Bits,
			&row.Compression, &row.Clientdn)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
