package pgstats

import (
	"errors"
	"strconv"
)

// Host sets the host to connect to.
// Values that start with / are for unix domain sockets.
//
// Default: localhost
func Host(host string) Option {
	return func(s *connection) error {
		s.config["host"] = host
		return nil
	}
}

// Port sets the port to bind to.
//
// Default: 5432
func Port(port int) Option {
	return func(s *connection) error {
		s.config["port"] = strconv.Itoa(port)
		return nil
	}
}

// SslMode sets whether or not to use SSL.
// Valid values for sslmode are:
//     disable - No SSL
//     require - Always SSL (skip verification)
//     verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
//     verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)
// Default: require
func SslMode(mode string) Option {
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

// FallbackApplicationName sets an application_name
// to fall back to if one isn't provided.
func FallbackApplicationName(name string) Option {
	return func(s *connection) error {
		s.config["fallback_application_name"] = name
		return nil
	}
}

// ConnectTimeout sets the maximum wait for connection, in seconds.
// Zero or not specified means wait indefinitely.
func ConnectTimeout(seconds int) Option {
	return func(s *connection) error {
		s.config["connect_timeout"] = strconv.Itoa(seconds)
		return nil
	}
}

// SslCert sets the location of the certificate file.
// The file must contain PEM encoded data.
func SslCert(location string) Option {
	return func(s *connection) error {
		s.config["sslcert"] = location
		return nil
	}
}

// SslKey sets the key file location.
// The file must contain PEM encoded data.
func SslKey(location string) Option {
	return func(s *connection) error {
		s.config["sslkey"] = location
		return nil
	}
}

// SslRootCert sets the location of the root certificate file.
// The file must contain PEM encoded data.
func SslRootCert(location string) Option {
	return func(s *connection) error {
		s.config["sslrootcert"] = location
		return nil
	}
}
