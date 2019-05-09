package pgstats

import "database/sql"

// PgStatUserFunctionsView represents content of pg_stat_user_functions view
type PgStatUserFunctionsView []PgStatFunctionsRow

// PgStatUserFunctionsView represents content of pg_stat_xact_user_functions view
type PgStatXactUserFunctionsView []PgStatFunctionsRow

// PgStatFunctionsRow represents schema of pg_stat*_user_functions views
type PgStatFunctionsRow struct {
	// OID of a function
	FuncId int64
	// Name of the schema this function is in
	SchemaName string
	// Name of this function
	FuncName string
	// Number of times this function has been called
	Calls sql.NullInt64
	// Total time spent in this function and all other functions called by it, in milliseconds
	TotalTime sql.NullFloat64
	// Total time spent in this function itself, not including other functions called by it, in milliseconds
	SelfTime sql.NullFloat64
}

func (s *PgStats) fetchFunctions(view string) ([]PgStatFunctionsRow, error) {
	data := make([]PgStatFunctionsRow, 0)
	db := s.conn.db
	query := "select funcid,schemaname,funcname,calls,total_time,self_time from " + view

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		row := new(PgStatFunctionsRow)
		err := rows.Scan(&row.FuncId, &row.SchemaName, &row.FuncName, &row.Calls, &row.TotalTime, &row.SelfTime)
		if err != nil {
			return nil, err
		}
		data = append(data, *row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
