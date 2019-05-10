package pgstats

import "database/sql"

// PgStatIoAllSequencesView represents content of pg_statio_all_sequences view
type PgStatIoAllSequencesView []PgStatIoSequencesRow

// PgStatIoSystemSequencesView represents content of pg_statio_system_sequences view
type PgStatIoSystemSequencesView []PgStatIoSequencesRow

// PgStatIoUserSequencesView represents content of pg_statio_user_sequences view
type PgStatIoUserSequencesView []PgStatIoSequencesRow

// PgStatIoSequencesRow represents schema of pg_statio_*_sequences views
type PgStatIoSequencesRow struct {
	// OID of a sequence
	RelId int64 `json:"relid"`
	// Name of the schema this sequence is in
	SchemaName string `json:"schemaname"`
	// Name of this sequence
	RelName string `json:"relname"`
	// Number of disk blocks read from this sequence
	BlksRead sql.NullInt64 `json:"blks_read"`
	// Number of buffer hits in this sequence
	BlksHit sql.NullInt64 `json:"blks_hit"`
}

func (s *PgStats) fetchIoSequences(view string) ([]PgStatIoSequencesRow, error) {
	db := s.conn.db
	query := "select relid,schemaname,relname,blks_read,blks_hit from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatIoSequencesRow, 0)
	for rows.Next() {
		row := new(PgStatIoSequencesRow)
		err := rows.Scan(&row.RelId, &row.SchemaName, &row.RelName, &row.BlksRead, &row.BlksHit)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
