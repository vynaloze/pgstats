// +build integration

package pgstats_test

import (
	"github.com/vynaloze/pgstats"
	"strings"
	"testing"
)

func TestReturnErrorOnUndefinedConnection(t *testing.T) {
	_, err := pgstats.PgStatActivity()
	if err == nil || err.Error() != "connection has not been defined" {
		t.Error("Wrong or no error returned")
	}
}

func TestDefineConnectionTwice(t *testing.T) {
	t.Parallel()
	errs := make(chan error, 2)
	def := func() {
		errs <- pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	}

	go def()
	go def()

	res := make([]error, 2)
	res[0] = <-errs
	res[1] = <-errs

	if res[0] != nil {
		t.Error(res[0])
	}
	if res[1] != nil {
		t.Error(res[1])
	}
}

func TestPgActivityWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	a, err := pgstats.PgStatActivity()
	validate(t, len(a), err)
}

func TestPgReplicationWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = pgstats.PgStatReplication()
	if err != nil {
		t.Error(err)
	}
}

func TestPgWalReceiverWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = pgstats.PgStatWalReceiver()
	if err != nil && !strings.Contains(err.Error(), "Unsupported PostgreSQL version: 9.") {
		if err.Error() != "sql: no rows in result set" { //fixme hack until there is test env for replication stats
			t.Error(err)
		}
	}
}

func TestPgSubscriptionWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = pgstats.PgStatSubscription()
	if err != nil && !strings.Contains(err.Error(), "Unsupported PostgreSQL version: 9.") {
		t.Error(err)
	}
}

func TestPgSslWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	a, err := pgstats.PgStatSsl()
	if err != nil {
		if strings.Contains(err.Error(), "Unsupported PostgreSQL version: 9.") {
			return
		}
		t.Error(err)
	}
	if len(a) == 0 {
		t.Error("No data returned by query")
	}
}

func TestPgStatProgressVacuumWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = pgstats.PgStatProgressVacuum()
	if err != nil && !strings.Contains(err.Error(), "Unsupported PostgreSQL version: 9.") {
		t.Error(err)
	}
}

func TestPgArchiverWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = pgstats.PgStatArchiver()
	if err != nil {
		t.Error(err)
	}
}

func TestPgBgWriterWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	_, err = pgstats.PgStatBgWriter()
	if err != nil {
		t.Error(err)
	}
}

func TestPgStatDatabaseWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	d, err := pgstats.PgStatDatabase()
	validate(t, len(d), err)
}

func TestPgStatDatabaseConflictsWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	d, err := pgstats.PgStatDatabaseConflicts()
	validate(t, len(d), err)
}

func TestPgStatTablesWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := pgstats.PgStatAllTables()
	validate(t, len(all), err)
	sys, err := pgstats.PgStatSystemTables()
	validate(t, len(sys), err)
	usr, err := pgstats.PgStatUserTables()
	validate(t, len(usr), err)
}

func TestPgXactStatTablesWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := pgstats.PgStatXactAllTables()
	validate(t, len(all), err)
	sys, err := pgstats.PgStatXactSystemTables()
	validate(t, len(sys), err)
	usr, err := pgstats.PgStatXactUserTables()
	validate(t, len(usr), err)
}

func TestPgStatIndexesWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := pgstats.PgStatAllIndexes()
	validate(t, len(all), err)
	sys, err := pgstats.PgStatSystemIndexes()
	validate(t, len(sys), err)
	usr, err := pgstats.PgStatUserIndexes()
	validate(t, len(usr), err)
}

func TestPgStatIoTablesWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := pgstats.PgStatIoAllTables()
	validate(t, len(all), err)
	sys, err := pgstats.PgStatIoSystemTables()
	validate(t, len(sys), err)
	usr, err := pgstats.PgStatIoUserTables()
	validate(t, len(usr), err)
}

func TestPgStatIoIndexesWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := pgstats.PgStatIoAllIndexes()
	validate(t, len(all), err)
	sys, err := pgstats.PgStatIoSystemIndexes()
	validate(t, len(sys), err)
	usr, err := pgstats.PgStatIoUserIndexes()
	validate(t, len(usr), err)
}

func TestPgStatIoSequencesWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	all, err := pgstats.PgStatIoAllSequences()
	validate(t, len(all), err)
	_, err = pgstats.PgStatIoSystemSequences()
	if err != nil {
		t.Error(err)
	}
	usr, err := pgstats.PgStatIoUserSequences()
	validate(t, len(usr), err)
}

func TestPgStatFunctionsWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	usr, err := pgstats.PgStatUserFunctions()
	validate(t, len(usr), err)
	_, err = pgstats.PgStatXactUserFunctions()
	if err != nil {
		t.Error(err)
	}
}

func TestPgStatStatementsWrapper(t *testing.T) {
	t.Parallel()
	err := pgstats.DefineConnection(*dbname, *user, *password, pgstats.SslMode("disable"))
	if err != nil {
		t.Error(err)
	}
	ss, err := pgstats.PgStatStatements()
	validate(t, len(ss), err)
}
