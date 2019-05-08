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

func TestPgStatTables(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatAllTables()
	if err != nil || len(all) == 0 {
		t.Error(err)
	}
	sys, err := s.PgStatSystemTables()
	if err != nil || len(sys) == 0 {
		t.Error(err)
	}
	usr, err := s.PgStatUserTables()
	if err != nil || len(usr) == 0 {
		t.Error(err)
	}
}

func TestPgStatIndexes(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatAllIndexes()
	if err != nil || len(all) == 0 {
		t.Error(err)
	}
	sys, err := s.PgStatSystemIndexes()
	if err != nil || len(sys) == 0 {
		t.Error(err)
	}
	usr, err := s.PgStatUserIndexes()
	if err != nil || len(usr) == 0 {
		t.Error(err)
	}
}

func TestPgStatIoTables(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatIoAllTables()
	if err != nil || len(all) == 0 {
		t.Error(err)
	}
	sys, err := s.PgStatIoSystemTables()
	if err != nil || len(sys) == 0 {
		t.Error(err)
	}
	usr, err := s.PgStatIoUserTables()
	if err != nil || len(usr) == 0 {
		t.Error(err)
	}
}

func TestPgStatIoIndexes(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatIoAllIndexes()
	if err != nil || len(all) == 0 {
		t.Error(err)
	}
	sys, err := s.PgStatIoSystemIndexes()
	if err != nil || len(sys) == 0 {
		t.Error(err)
	}
	usr, err := s.PgStatIoUserIndexes()
	if err != nil || len(usr) == 0 {
		t.Error(err)
	}
}
