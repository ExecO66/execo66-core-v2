# Execo66 Core API Service

Core service for Execo66, a code assignment auto-grader.

## Quick Start

```shell
git clone https://github.com/ExecO66/execo66-core-v2.git

cd execo66-core-v2
```

For the following installations, you must have [Go 1.17+](https://go.dev/doc/install).

### Required Installations:

- **Make** for [MacOS](https://formulae.brew.sh/formula/make) or [Windows](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows) developers
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) - migration CLI tool
- [Postgres](https://www.postgresql.org/download) - database
- [Air](https://github.com/cosmtrek/air) - Live reload

### Database setup

Run the following commands

```shell
psql -U postgres
> CREATE DATABASE execo;
> \c execo
```

### Config Files

Create ./config/.env.dev file
Follow the .env.example for setup:

```none
PSQL_CONN_STRING=postgres://postgres:postgres@localhost:5432/execo?sslmode=disable
COOKIE_SECRET=secret
```

Run in development environment:

```shell
make migrate
make seed-db
make run-dev
```

## Project Layout

Execo Core uses the following project layout:

```none
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── config               configuration files
├── internal             private application and library code
│   ├── api              routes and handlers
│   ├── config           configuration variables
│   ├── entity           entity definitions
│   │   └─ queries       db queries
│   └── test             helpers for testing purpose
└── db                   database scripts
    ├── migrations       database migrations
    └── seeding          database test data
```
