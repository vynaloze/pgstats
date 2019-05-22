package pgstats

import (
	"github.com/pkg/errors"
	"regexp"
)

func (s *PgStats) getPgVersion() (string, error) {
	db := s.conn.db
	query := "show server_version;"
	row := db.QueryRow(query)
	version := new(string)
	err := row.Scan(&version)
	if err != nil {
		return "", err
	}
	v := extractMajorVersion(*version)
	if v == "" {
		err = errors.New("Regex error")
	}
	return v, err
}

func extractMajorVersion(fullVersion string) string {
	r, err := regexp.Compile(`(^9\.\d)|(^\d\d)`)
	if err != nil {
		return ""
	}
	return r.FindString(fullVersion)
}
