# go_blog

 A blog project with go.

## build and run

```sh
go build main.go && ./main.exe
```

## run with docker-compose

### start mysql and redis

if docker compose v1

```sh
docker-compose -f docker-mysql.yml -f docker-redis.yml up -d
```

if docker compose v2

```sh
docker-compose -p go_blog_db -f docker-mysql.yml -f docker-redis.yml up -d
```

stop if v1

```sh
docker-compose -f docker-mysql.yml -f docker-redis.yml down
```

stop if v2

```sh
docker-compose -p go_blog_db down
```

### start go_blog app

start

```sh
docker-compose up -d
```

stop

```sh
docker-compose down
```

## create database

create database with collate 'utf8mb4_unicode_ci'

```sql
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

```sh
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
