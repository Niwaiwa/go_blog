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
