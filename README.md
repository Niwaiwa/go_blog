# go_blog

 A blog project with go.


## run with docker-compose

### start mysql and redis

if docker compose v1

```
docker-compose -f docker-mysql.yml -f docker-redis.yml up -d
```

if docker compose v2

```
docker-compose -p go_blog_db -f docker-mysql.yml -f docker-redis.yml up -d
```

stop if v1

```
docker-compose -f docker-mysql.yml -f docker-redis.yml down
```

stop if v2

```
docker-compose -p go_blog_db down
```

### start go_blog app

```
docker-compose up -d
```

stop

```
docker-compose down
```

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
