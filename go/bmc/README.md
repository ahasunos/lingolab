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

## Using TablePlus for interacting with DB
- https://tableplus.com/
