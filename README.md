# Execo66 Core API Service

Core service for Execo66, a code assignment auto-grader.

## Quick Start

```shell
git clone https://github.com/ExecO66/execo66-core-v2.git

cd execo66-core-v2
```

For the following installations, you must have [Go 1.17+](https://go.dev/doc/install).

### Required Installations:

- [Task](https://taskfile.dev/#/installation?id=installation) - Task runner and build tool
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) - migration CLI tool
- [Postgres](https://www.postgresql.org/download) - database
- [Air](https://github.com/cosmtrek/air) - Live reload (not required, nice to have)

### Database setup

Run the following commands

```shell
psql -U postgres
> CREATE DATABASE execo;
> CREATE DATABASE execo_testing; # required when running query tests
> \c execo
```

### Config Files

Create ./config/.env.dev file
Follow the .env.example for setup:

<sup>\*\*only TEST_PSQL_CONN_STRING is required when working with query tests</sup>

```none
PSQL_CONN_STRING=postgres://postgres:postgres@localhost:5432/execo?sslmode=disable
TEST_PSQL_CONN_STRING=postgres://postgres:postgres@localhost:5432/execo_testing?sslmode=disable
COOKIE_SECRET=secret
PORT=8080
GOOGLE_AUTH_CLIENT_ID=123abc
GOOGLE_AUTH_CLIENT_SECRET=456def
GOOGLE_AUTH_REDIRECT_URI=http://localhost:8080/auth/google/callback
CLIENT_BASE_URL=http://localhost:3000
```

#### Run API in development environment

```shell
task migrate
task seed-db
task run-dev
```

#### Run query tests

```shell
task test-queries-all
```

or if you want more control:

```shell
task migrate-t
task seed-t-db
task test-queries
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
