# PgConn

pgconncheck is a simple postgresql that will run a connection test and log any issues or disturbances.

## Running

### Options

Options can be passed via the command line or via environment variables.

```bash
pgconncheck -h 

  -pgdb string
        [PGDB] Database name (default "test")
  -pghost string
        [PGHOST] Database host (default "localhost")
  -pgpassword string
        [PGPASSWORD] Database password (default "PgConn")
  -pgport int
        [PGPORT] Database master port (default 5432)
  -pgsleep int
        [PGSLEEP] Time to sleep between checks (default 100)
  -pguser string
        [PGUSER] Database user (default "postgres")
```

### Testing

```bash
docker run  --rm --name pgconn -e "POSTGRES_PASSWORD=PgConn" -e "POSTGRES_DB=test" -p 5432:5432 postgres
```

Then kill and restart the container.

### Sample Output

```bash
2024/06/10 14:58:36 Connecting to postgres://postgres:XXXXX@localhost:5432/test
2024/06/10 14:58:54 Error: driver: bad connection postgres://postgres:XXXXX@localhost:5432/test
2024/06/10 14:58:56 Connected to postgres://postgres:XXXXX@localhost:5432/test
```
