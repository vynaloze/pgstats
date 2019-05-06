package pgstats

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
)

type connectionConfig map[string]string

type connection struct {
	config     connectionConfig
	connString string
	db         *sql.DB
}

func (s *PgStats) prepareConnection(dbname string, user string, password string, options ...func(*connection) error) error {
	conn := &connection{}
	conn.setRequiredParams(dbname, user, password)
	err := conn.setOptionalParams(options...)
	if err != nil {
		return err
	}
	conn.buildConnectionString()
	s.conn = conn
	return nil
}

func (s *PgStats) openConnection() error {
	// Create connection
	db, err := sql.Open("postgres", s.conn.connString)
	if err != nil {
		return err
	}

	// Open connection
	err = db.Ping()
	if err != nil {
		return err
	}

	s.conn.db = db
	return nil
}

func (c *connection) setRequiredParams(dbname string, user string, password string) {
	config := make(connectionConfig)
	config["dbname"] = dbname
	config["user"] = user
	config["password"] = password
	c.config = config
}

func (c *connection) setOptionalParams(options ...func(*connection) error) error {
	for _, option := range options {
		err := option(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *connection) buildConnectionString() {
	var str strings.Builder
	for param, value := range c.config {
		str.WriteString(param)
		str.WriteString("=")
		str.WriteString(value)
		str.WriteString(" ")
	}
	c.connString = str.String()
}
