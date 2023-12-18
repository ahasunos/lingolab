# BMC (Backend Master Class)

## Database Design
- Design a SQL DB schema using [dbdiagram.io](https://dbdiagram.io/)
- [Diagram Link](https://dbdiagram.io/d/SimpleBank-657ff2e356d8064ca0366da0)
![DB Diagram](SimpleBank.png)

## Using Postgres from Docker
- Link to postgres image: https://hub.docker.com/_/postgres
- Pull the image: `docker pull postgres:16-alpine`
- Run the image: `docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:16-alpine`
- Exec in the container: `docker exec -it postgres16 psql -U root`
  ```
  ❯ docker exec -it postgres16 psql -U root
  psql (16.1)
  Type "help" for help.

  root=# select now();
                now
  -------------------------------
  2023-12-18 07:49:57.791533+00
  (1 row)

  root=# \q
  ```
- View the logs: `docker logs postgres16`
- To stop the container: `docker stop postgres16`

## Using TablePlus for interacting with DB
- https://tableplus.com/

## DB Migration
- Use golang migrate: https://github.com/golang-migrate/migrate
- Documentation: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Usage

  ```
  $ migrate -help
  Usage: migrate OPTIONS COMMAND [arg...]
        migrate [ -version | -help ]

  Options:
    -source          Location of the migrations (driver://url)
    -path            Shorthand for -source=file://path
    -database        Run migrations against this database (driver://url)
    -prefetch N      Number of migrations to load in advance before executing (default 10)
    -lock-timeout N  Allow N seconds to acquire database lock (default 15)
    -verbose         Print verbose logging
    -version         Print version
    -help            Print usage

  Commands:
    create [-ext E] [-dir D] [-seq] [-digits N] [-format] NAME
                Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
                Use -seq option to generate sequential up/down migrations with N digits.
                Use -format option to specify a Go time format string.
    goto V       Migrate to version V
    up [N]       Apply all or N up migrations
    down [N]     Apply all or N down migrations
    drop         Drop everything inside database
    force V      Set version V but don't run migration (ignores dirty state)
    version      Print current migration version
  ```
- Create the migration file: `migrate create -ext sql -dir db/migration -seq init_schema`

  ```
  ❯ migrate create -ext sql -dir db/migration -seq init_schema
  ****/lingolab/lingolab/go/bmc/db/migration/000001_init_schema.up.sql
  ****/lingolab/lingolab/go/bmc/db/migration/000001_init_schema.down.sql
  ```
- Up script is used to run the forward change to the schema and down script is used to revert the change
- Perform the migration: `migrate -path db/migration/ -database "postgresql://root:mysecret@localhost:5432/simple_bank" -verbose up`

  ```
  ❯ migrate -path db/migration/ -database "postgresql://root:mysecret@localhost:5432/simple_bank" -verbose up
  2023/12/18 16:08:29 error: pq: SSL is not enabled on the server
  ```

  By default SSL is not enabled to the postgres server, so pass `?sslmode=disable`
  ```
  ❯ migrate -path db/migration/ -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose up
  2023/12/18 16:13:15 Start buffering 1/u init_schema
  2023/12/18 16:13:15 Read and execute 1/u init_schema
  2023/12/18 16:13:15 Finished 1/u init_schema (read 7.259083ms, ran 31.139959ms)
  2023/12/18 16:13:15 Finished after 44.083916ms
  2023/12/18 16:13:15 Closing source and database
  ```

## Create and Drop DB in Postgres
```
❯ docker stop postgres16
postgres16

❯ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

❯ docker ps -a
CONTAINER ID   IMAGE                                   COMMAND                  CREATED         STATUS                      PORTS     NAMES
50cdc9d7aabb   postgres:16-alpine                      "docker-entrypoint.s…"   3 hours ago     Exited (0) 30 seconds ago             postgres16

❯ docker start postgres16
postgres16

❯ docker ps
CONTAINER ID   IMAGE                COMMAND                  CREATED       STATUS        PORTS                    NAMES
50cdc9d7aabb   postgres:16-alpine   "docker-entrypoint.s…"   3 hours ago   Up 1 second   0.0.0.0:5432->5432/tcp   postgres16

❯ docker exec -it postgres16 /bin/sh
/ # ls
bin                         etc                         media                       proc                        sbin                        tmp
dev                         home                        mnt                         root                        srv                         usr
docker-entrypoint-initdb.d  lib                         opt                         run                         sys                         var
/ # createdb --username=root --owner=root simple_bank
/ # psql simple_bank
psql (16.1)
Type "help" for help.

simple_bank=# select now();
              now
-------------------------------
 2023-12-18 10:24:07.360775+00
(1 row)

simple_bank=# \q
/ # dropdb simple_bank
/ # psql simple_bank
psql: error: connection to server on socket "/var/run/postgresql/.s.PGSQL.5432" failed: FATAL:  database "simple_bank" does not exist
/ # exit
```

### Create without getting in the container
```
❯ docker exec -it postgres16 createdb --username=root --owner=root simple_bank

❯ docker exec -it postgres16 psql -U root simple_bank
psql (16.1)
Type "help" for help.

simple_bank=# exit
```

### Using Makefile
```
❯ make postgres
docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:16-alpine
docker: Error response from daemon: Conflict. The container name "/postgres16" is already in use by container "50cdc9d7aabb2bd17639c4794a75354e92b0a3c43c173edb9faa147ac4d2b259". You have to remove (or rename) that container to be able to reuse that name.
See 'docker run --help'.
make: *** [postgres] Error 125

❯ docker rm postgres16
postgres16

❯ make postgres
docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:16-alpine
17d03307706a98a4550203c57a493e45ed8bacae974ab656df2091e457eed8b4

❯ docker ps
CONTAINER ID   IMAGE                COMMAND                  CREATED         STATUS         PORTS                    NAMES
17d03307706a   postgres:16-alpine   "docker-entrypoint.s…"   2 seconds ago   Up 2 seconds   0.0.0.0:5432->5432/tcp   postgres16

❯ make createdb
docker exec -it postgres16 createdb --username=root --owner=root simple_bank

❯ make migratedown
migrate -path db/migration/ -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose down
2023/12/18 16:17:33 Are you sure you want to apply all down migrations? [y/N]
y
2023/12/18 16:17:36 Applying all down migrations
2023/12/18 16:17:36 Start buffering 1/d init_schema
2023/12/18 16:17:36 Read and execute 1/d init_schema
2023/12/18 16:17:36 Finished 1/d init_schema (read 27.455542ms, ran 27.595583ms)
2023/12/18 16:17:36 Finished after 2.506294542s
2023/12/18 16:17:36 Closing source and database

❯ make migrateup
migrate -path db/migration/ -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" -verbose up
2023/12/18 16:18:04 Start buffering 1/u init_schema
2023/12/18 16:18:04 Read and execute 1/u init_schema
2023/12/18 16:18:04 Finished 1/u init_schema (read 8.353375ms, ran 31.413084ms)
2023/12/18 16:18:04 Finished after 45.6405ms
2023/12/18 16:18:04 Closing source and database
```
