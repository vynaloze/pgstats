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

func TestPgActivity(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	a, err := s.PgStatActivity()
	validate(t, len(a), err)
}

func TestPgReplication(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = s.PgStatReplication()
	if err != nil {
		t.Error(err)
	}
}

func TestPgWalReceiver(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = s.PgStatWalReceiver()
	if err != nil {
		if err.Error() != "sql: no rows in result set" { //fixme hack until there is test env for replication stats
			t.Error(err)
		}
	}
}

func TestPgSubscription(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = s.PgStatSubscription()
	if err != nil {
		t.Error(err)
	}
}

func TestPgArchiver(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = s.PgStatArchiver()
	if err != nil {
		t.Error(err)
	}
}

func TestPgBgWriter(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = s.PgStatBgWriter()
	if err != nil {
		t.Error(err)
	}
}

func TestPgStatDatabase(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	d, err := s.PgStatDatabase()
	validate(t, len(d), err)
}

func TestPgStatDatabaseConflicts(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	d, err := s.PgStatDatabaseConflicts()
	validate(t, len(d), err)
}

func TestPgStatTables(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatAllTables()
	validate(t, len(all), err)
	sys, err := s.PgStatSystemTables()
	validate(t, len(sys), err)
	usr, err := s.PgStatUserTables()
	validate(t, len(usr), err)
}

func TestPgXactStatTables(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatXactAllTables()
	validate(t, len(all), err)
	sys, err := s.PgStatXactSystemTables()
	validate(t, len(sys), err)
	usr, err := s.PgStatXactUserTables()
	validate(t, len(usr), err)
}

func TestPgStatIndexes(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatAllIndexes()
	validate(t, len(all), err)
	sys, err := s.PgStatSystemIndexes()
	validate(t, len(sys), err)
	usr, err := s.PgStatUserIndexes()
	validate(t, len(usr), err)
}

func TestPgStatIoTables(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatIoAllTables()
	validate(t, len(all), err)
	sys, err := s.PgStatIoSystemTables()
	validate(t, len(sys), err)
	usr, err := s.PgStatIoUserTables()
	validate(t, len(usr), err)
}

func TestPgStatIoIndexes(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatIoAllIndexes()
	validate(t, len(all), err)
	sys, err := s.PgStatIoSystemIndexes()
	validate(t, len(sys), err)
	usr, err := s.PgStatIoUserIndexes()
	validate(t, len(usr), err)
}

func TestPgStatIoSequences(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := s.PgStatIoAllSequences()
	validate(t, len(all), err)
	_, err = s.PgStatIoSystemSequences()
	if err != nil {
		t.Error(err)
	}
	usr, err := s.PgStatIoUserSequences()
	validate(t, len(usr), err)
}

func TestPgStatFunctions(t *testing.T) {
	s, err := pgstats.Connect(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	usr, err := s.PgStatUserFunctions()
	validate(t, len(usr), err)
	_, err = s.PgStatXactUserFunctions()
	if err != nil {
		t.Error(err)
	}
}

func validate(t *testing.T, len int, err error) {
	if err != nil {
		t.Error(err)
	}
	if len == 0 {
		t.Error("No data returned by query")
	}
}
