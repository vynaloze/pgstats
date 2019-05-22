package pgstats

import (
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

func (s *PgStats) getPgVersion() (float64, error) {
	db := s.conn.db
	query := "show server_version;"
	row := db.QueryRow(query)
	version := new(string)
	if err := row.Scan(&version); err != nil {
		return 0, err
	}
	return extractMajorVersion(*version)
}

func extractMajorVersion(fullVersion string) (float64, error) {
	r, err := regexp.Compile(`(^9\.\d)|(^\d\d)`)
	if err != nil {
		return 0, errors.New("Regex compile error")
	}
	strv := r.FindString(fullVersion)
	if strv == "" {
		return 0, errors.New("Regex parse error")
	}
	version, err := strconv.ParseFloat(strv, 64)
	if err != nil {
		return 0, err
	}
	return version, nil
}
