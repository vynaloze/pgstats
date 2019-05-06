package pgstats

import (
	"errors"
	"strconv"
)

func Host(host string) func(*connection) error {
	return func(s *connection) error {
		s.config["host"] = host
		return nil
	}
}

func Port(port int) func(*connection) error {
	return func(s *connection) error {
		s.config["port"] = strconv.Itoa(port)
		return nil
	}
}

func SslMode(mode string) func(*connection) error {
	validModes := map[string]struct{}{
		"disable":     {},
		"require":     {},
		"verify-ca":   {},
		"verify-full": {},
	}
	return func(s *connection) error {
		if _, ok := validModes[mode]; ok {
			s.config["sslmode"] = mode
			return nil
		}
		return errors.New("Invalid SSL mode. Allowed values: disable, require, verify-ca, verify-full")
	}
}
