# go_blog

 A blog project with go.

## create mysql database

start database container with docker

```
docker-compose -f docker-mysql.yml up -d
```

create database with collate 'utf8mb4_unicode_ci'

```
CREATE DATABASE blog COLLATE 'utf8mb4_unicode_ci';
```

## migration with golang-migrate/migrate

download migration file

* [golang-migrate/migrate](https://github.com/golang-migrate/migrate/releases)

mysql connection

```text
mysql://user:password@tcp(host:port)/dbname
mysql://root:password@tcp(127.0.0.1:3306)/blog
```

create migration file

```
migrate create -ext sql -dir db/migrations -seq create_users
```

migrate up

```sh
migrate -database ${DATABASE_URL} -path db/migrations up 1
```

migrate down

```sh
migrate -database ${DATABASE_URL} -path db/migrations down 1
```

migrate force

```sh
migrate -database ${DATABASE_URL} -path db/migrations force 1
```
