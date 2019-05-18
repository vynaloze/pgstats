## Optional connection parameters
Inspired by [this post](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis).

### Usage
Specify zero or more optional parameters in a form of a function
```go
err := pgstats.DefineConnection("foo", "username", "password")
// or
conn, err := pgstats.Connect("foo", "username", "password")
```
```go
err := pgstats.DefineConnection("foo", "username", "password", pgstats.SslMode("disable"))
// or
conn, err := pgstats.Connect("foo", "username", "password", pgstats.SslMode("disable"))
```
```go
err := pgstats.DefineConnection("foo", "username", "password", Host("10.0.1.3"), Port(6432), 
	SslMode("require-ca"), FallbackApplicationName("testApp"), ConnectTimeout(42), 
	SslCert("/test/loc"), SslKey("/another"), SslRootCert("/i/am/already/tired/of/this"))
// or
conn, err := pgstats.Connect("foo", "username", "password", Host("10.0.1.3"), Port(6432), 
	SslMode("require-ca"), FallbackApplicationName("testApp"), ConnectTimeout(42), 
	SslCert("/test/loc"), SslKey("/another"), SslRootCert("/i/am/already/tired/of/this"))
   ```

### Overview

- [ConnectTimeout](#func--connecttimeout)
- [FallbackApplicationName](#func--fallbackapplicationname)
- [Host](#func--host)
- [Port](#func--port)
- [SslCert](#func--sslcert)
- [SslKey](#func--sslkey)
- [SslMode](#func--sslmode)
- [SslRootCert](#func--sslrootcert)

#### func  ConnectTimeout

```go
func ConnectTimeout(seconds int) func(*connection) error
```
ConnectTimeout sets the maximum wait for connection, in seconds. Zero or not
specified means wait indefinitely.

#### func  FallbackApplicationName

```go
func FallbackApplicationName(name string) func(*connection) error
```
FallbackApplicationName sets an application_name to fall back to if one isn't
provided.

#### func  Host

```go
func Host(host string) func(*connection) error
```
Host sets the host to connect to. Values that start with / are for unix domain
sockets.

Default: localhost

#### func  Port

```go
func Port(port int) func(*connection) error
```
Port sets the port to bind to.

Default: 5432

#### func  SslCert

```go
func SslCert(location string) func(*connection) error
```
SslCert sets the location of the certificate file. The file must contain PEM
encoded data.

#### func  SslKey

```go
func SslKey(location string) func(*connection) error
```
SslKey sets the key file location. The file must contain PEM encoded data.

#### func  SslMode

```go
func SslMode(mode string) func(*connection) error
```
SslMode sets whether or not to use SSL. Valid values for sslmode are:

    disable - No SSL
    require - Always SSL (skip verification)
    verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
    verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)

Default: require

#### func  SslRootCert

```go
func SslRootCert(location string) func(*connection) error
```
SslRootCert sets the location of the root certificate file. The file must
contain PEM encoded data.
