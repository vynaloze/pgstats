# pgstats
[![Documentation](https://godoc.org/github.com/vynaloze/pgstats?status.svg)](https://godoc.org/github.com/vynaloze/pgstats)
[![Go Report Card](https://goreportcard.com/badge/github.com/vynaloze/pgstats)](https://goreportcard.com/report/github.com/vynaloze/pgstats)
[![Coverage Status](https://coveralls.io/repos/github/vynaloze/pgstats/badge.svg?branch=master)](https://coveralls.io/github/vynaloze/pgstats?branch=master)
[![Build Status](https://travis-ci.com/vynaloze/pgstats.svg?branch=master)](https://travis-ci.com/vynaloze/pgstats)

**pgstats** provides convenient access to **pg_stat_&ast;** statistics, allowing to monitor **PostgreSQL** instances inside **go** applications.

## Attention - Work in Progress!
The library is in _early access_ state, under active development. 
Expect breaking changes, new functionalities and improvements in the near future. 
Use at your own risk.

## Install
`go get github.com/vynaloze/pgstats`

## API reference
Check out the [quick API overview](https://github.com/vynaloze/pgstats/wiki/API-methods) or [full documentation on godoc.org](https://godoc.org/github.com/vynaloze/pgstats)

## Usage
### Want it simple?
**It will be supported since v1.0.0!**

1. **Define your connection.** You can do it anywhere, at any time and as many times as you want.
However, you cannot override the settings once you define the connection. If you want to play with many connections, 
see the [next section](#want-to-have-multiple-connections)

    ```go
    err := pgstats.DefineConnection("foo", "username", "password")
    ```
    
2. **Now you can collect statistics in any part of your code.** 
If the connection has not been defined before, an error is returned.

    ```go
    // pg_stat_bgwriter - returns single row
    b, _ := pgstats.PgStatBgWriter()
    fmt.Println(b.CheckpointsTimed)
    // Example result:
    // {446 true}
    
    // pg_stat_user_tables - returns many rows
    uts, _ := pgstats.PgStatUserTables()
    for _, ut := range uts {
       fmt.Printf("%s - seq_tup_read: %v\n", ut.Relname, ut.SeqTupRead.Int64)
    }
    // Example result:
    // foo - seq_tup_read: 9273
    // bar - seq_tup_read: 10
    ```
    
    
### Want to have multiple connections?
1. **Define them.** If you want to free the connection pool after you are done 
(and you are not exiting your application right away), 
you can close the connection with Close() method.

    ```go
    connFoo, _ := pgstats.Connect("foo", "username", "password")
    defer connFoo.Close()
    connBar, _ := pgstats.Connect("bar", "username", "password")
    defer connBar.Close()
    ```
    
2. **Use them.**

    ```go
    // Query both connections
    utf, _ := connFoo.PgStatUserTables()
    utb, _ := connBar.PgStatUserTables()
    // Print first entries in pg_stat_user_tables for both databases
    fmt.Printf("foo: %s - seq_tup_read: %v\n", utf[0].Relname, utf[0].SeqTupRead)
    fmt.Printf("bar: %s - seq_tup_read: %v\n", utb[0].Relname, utb[0].SeqTupRead)
    // foo: example - seq_tup_read: 9273
    // bar: test - seq_tup_read: 10
    ```

### Want to specify optional connection parameters?
No problem - use _functional options:_
```go
err := pgstats.DefineConnection("foo", "username", "password", pgstats.Host("10.0.1.3"), pgstats.Port(6432))
```
```go
conn, err := pgstats.Connect("foo", "username", "password", pgstats.SslMode("disable"))
```
[Full reference](https://github.com/vynaloze/pgstats/wiki/Connection-parameters)
## Supported PostgreSQL versions
- 11
- 10
- 9.6
- 9.5
- 9.4

[Full pg_stat view / version table](https://github.com/vynaloze/pgstats/wiki/Supported-stats-vs-PG-version)

## Roadmap
Done? | Version | Content | ETA
--- | --- | --- | ---
Y | `v0.1.0` | most stats available and tested in PG 10 | ~May '19
Y | | setup CI for integration tests | ~May '19
Y | `v0.2.0` | pg_stat_statements | ~June '19
Y | `v0.3.0` | support PG 11 | ~June '19
Y | `v0.4.0` | support PG 9.6 | ~June '19
Y | `v0.5.0` | support PG 9.5 | ~June '19
Y | `v0.6.0` | support PG 9.4 | ~June '19
N | `v1.0.0` | global wrapper for the ease of use | ~July '19
N | `v1.1.0` | add missing integration tests | ~July '19

## License
The library is licensed under the [MIT License](LICENSE).