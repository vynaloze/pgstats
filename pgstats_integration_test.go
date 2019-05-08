// +build integration

package pgstats_test

import (
	"flag"
	"github.com/vynaloze/pgstats"
	"testing"
)

var dbname = flag.String("dbname", "", "Test database name")
var user = flag.String("user", "", "Test username")
var password = flag.String("password", "", "Test password")

func TestPgStatIndexes(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatAllIndexes()
	if err != nil || len(all) == 0 {
		t.Error(err)
	}
	usr, err := s.PgStatUserIndexes()
	if err != nil || len(usr) == 0 {
		t.Error(err)
	}
	sys, err := s.PgStatSystemIndexes()
	if err != nil || len(sys) == 0 {
		t.Error(err)
	}
}
