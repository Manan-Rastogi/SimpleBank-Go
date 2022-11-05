# Set up Postgres using DOCKER
```
1. docker pull postgres:15.0-alpine

2. docker run --name postgres -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -d cc994726847f

3. docker exec -it postgres psql -U user

4. Check ---- select now(); --- \q to exit psql container

5. docker logs postgres

```

# Golang Migration
```
https://github.com/golang-migrate/migrate

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage
```

#   Migrate
## Changes from Go to DB and from DB to Go.
```
>   migrate -version
>   migrate --help
>   mkdir db/migration, then migrate create -ext sql -dir db/migration -seq init_schema
>   migrating up : migrate -path migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

```

# SQLC
### https://github.com/kyleconroy/sqlc
```
>   go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
>   sqlc version
>   sqlc --help
>   sqlc init : creates an yaml file.
>   go to v1.4 and go to settings, copy and paste under package.
>   Create query in query folder
>   refer to db/sqlc.yaml -> sqlc generate
>   https://github.com/kyleconroy/sqlc/tree/v1.4.0#getting-started
>   https://docs.sqlc.dev/en/latest/overview/install.html#docker   -- For using postgresql windows.

```
# Setting SQLC using docker on windows
```
>   docker pull kjconroy/sqlc
>   docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc version
>   docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc init
>   docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```


# Go Test File
```
>   Conventionally Test file name ends with test prefix.
>   Function Name starts with *Test*FUNCTION.
```


# ACID
### Transactions - A single unit of work, often made up of multiple db operations.
```
> Atomicity: Either all operations complete successfully or the transaction fails and the db is unchanged.
> Consistency: The db state must be valid after transaction.All constraints must be satisfied.
> Isolation: Concurrent transactions must not affect each other.
> Durability: Data written by successful transaction must be recorded in persistent storage.
```