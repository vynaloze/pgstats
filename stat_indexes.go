package pgstats

import (
	"github.com/vynaloze/pgstats/nullable"
)

// PgStatAllIndexesView represents content of pg_stat_all_indexes view
type PgStatAllIndexesView []PgStatIndexesRow

// PgStatSystemIndexesView represents content of pg_stat_system_indexes view
type PgStatSystemIndexesView []PgStatIndexesRow

// PgStatUserIndexesView represents content of pg_stat_user_indexes view
type PgStatUserIndexesView []PgStatIndexesRow

// PgStatIndexesRow represents schema of pg_stat_*_indexes views
type PgStatIndexesRow struct {
	// OID of the table for this index
	Relid int64 `json:"relid"`
	// OID of this index
	Indexrelid int64 `json:"indexrelid"`
	// Name of the schema this index is in
	Schemaname string `json:"schemaname"`
	// Name of the table for this index
	Relname string `json:"relname"`
	// Name of this index
	Indexrelname string `json:"indexrelname"`
	// Number of index scans initiated on this index
	IdxScan nullable.Int64 `json:"idx_scan"`
	// Number of index entries returned by scans on this index
	IdxTupRead nullable.Int64 `json:"idx_tup_read"`
	// Number of live table rows fetched by simple index scans using this index
	IdxTupFetch nullable.Int64 `json:"idx_tup_fetch"`
}

func (s *PgStats) fetchIndexes(view string) ([]PgStatIndexesRow, error) {
	db := s.conn.db
	query := "select relid,indexrelid,schemaname,relname,indexrelname," +
		"idx_scan,idx_tup_read,idx_tup_fetch from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatIndexesRow, 0)
	for rows.Next() {
		row := new(PgStatIndexesRow)
		err := rows.Scan(&row.Relid, &row.Indexrelid, &row.Schemaname, &row.Relname, &row.Indexrelname,
			&row.IdxScan, &row.IdxTupRead, &row.IdxTupFetch)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
