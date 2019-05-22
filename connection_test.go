package pgstats

import (
	"strings"
	"testing"
)

var connCorrectTestTable = []struct {
	dbname   string
	user     string
	password string
	options  []func(*connection) error
	expected []string
	count    int
}{
	{
		"testDb", "testUser", "12345",
		[]func(*connection) error{},
		[]string{"dbname=testDb", "user=testUser", "password=12345"},
		3,
	},
	{
		"testDb", "testUser", "12345",
		[]func(*connection) error{Host("testHostname")},
		[]string{"dbname=testDb", "user=testUser", "password=12345", "host=testHostname"},
		4,
	},
	{
		"testDb", "testUser", "12345",
		[]func(*connection) error{Host("testHostname"), Port(1234), SslMode("disable")},
		[]string{"dbname=testDb", "user=testUser", "password=12345", "host=testHostname", "port=1234", "sslmode=disable"},
		6,
	},
	{
		"testDb", "testUser", "12345",
		[]func(*connection) error{Host("testHostname"), Port(1234), SslMode("disable"),
			FallbackApplicationName("testApp"), ConnectTimeout(42), SslCert("/test/loc"),
			SslKey("/another"), SslRootCert("/i/am/already/tired/of/this")},
		[]string{"dbname=testDb", "user=testUser", "password=12345", "host=testHostname", "port=1234", "sslmode=disable",
			"fallback_application_name=testApp", "connect_timeout=42", "sslcert=/test/loc", "sslkey=/another", "sslrootcert=/i/am/already/tired/of/this"},
		11,
	},
}

var connFailTestTable = []struct {
	dbname   string
	user     string
	password string
	options  []func(*connection) error
}{
	{
		"testDb", "testUser", "12345",
		[]func(*connection) error{Host("testHostname"), Port(1234), SslMode("INVALID")},
	},
}

func TestCorrectConnection(t *testing.T) {
	t.Parallel()
	for _, tt := range connCorrectTestTable {
		s := PgStats{}
		err := s.prepareConnection(tt.dbname, tt.user, tt.password, tt.options...)
		if err != nil {
			t.Errorf("Unexpected error occurred: %s", err)
		} else {
			actual := s.conn.connString
			count := strings.Count(actual, "=")
			if count != tt.count {
				t.Errorf("Expected %d params; got %d", tt.count, count)
			}
			for _, expected := range tt.expected {
				if !strings.Contains(actual, expected) {
					t.Errorf("Expected '%s'; actual '%s'", tt.expected, actual)
				}
			}

		}
	}
}

func TestFailConnection(t *testing.T) {
	t.Parallel()
	for _, tt := range connFailTestTable {
		s := PgStats{}
		err := s.prepareConnection(tt.dbname, tt.user, tt.password, tt.options...)
		if err == nil {
			actual := s.conn.connString
			t.Errorf("Expected error, but it didn't happened. Actual data: %s", actual)
		}
	}
}
