package pgstats

import (
	"log"
)

type PgStats struct {
	conn *connection
}

func Connect(dbname string, user string, password string, options ...func(*connection) error) (*PgStats, error) {
	s := &PgStats{}
	err := s.prepareConnection(dbname, user, password, options...)
	if err != nil {
		log.Fatal(err)
	}
	err = s.openConnection()
	return s, err
}
