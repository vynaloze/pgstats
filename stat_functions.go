package pgstats

import "database/sql"

// PgStatUserFunctionsView represents content of pg_stat_user_functions view
type PgStatUserFunctionsView []PgStatFunctionsRow

// PgStatXactUserFunctionsView represents content of pg_stat_xact_user_functions view
type PgStatXactUserFunctionsView []PgStatFunctionsRow

// PgStatFunctionsRow represents schema of pg_stat*_user_functions views
type PgStatFunctionsRow struct {
	// OID of a function
	Funcid int64 `json:"funcid"`
	// Name of the schema this function is in
	Schemaname string `json:"schemaname"`
	// Name of this function
	Funcname string `json:"funcname"`
	// Number of times this function has been called
	Calls sql.NullInt64 `json:"calls"`
	// Total time spent in this function and all other functions called by it, in milliseconds
	TotalTime sql.NullFloat64 `json:"total_time"`
	// Total time spent in this function itself, not including other functions called by it, in milliseconds
	SelfTime sql.NullFloat64 `json:"self_time"`
}

func (s *PgStats) fetchFunctions(view string) ([]PgStatFunctionsRow, error) {
	db := s.conn.db
	query := "select funcid,schemaname,funcname,calls,total_time,self_time from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]PgStatFunctionsRow, 0)
	for rows.Next() {
		row := new(PgStatFunctionsRow)
		err := rows.Scan(&row.Funcid, &row.Schemaname, &row.Funcname, &row.Calls, &row.TotalTime, &row.SelfTime)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}
	return data, rows.Err()
}
