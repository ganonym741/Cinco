# CINCO
### CREATE DATABASEUSING POSTGRES

> Create database "**cinco**"

### INSTAL GOLANG MIGRATE
```shell
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```

### RUN Makefile
Create two files with up and down at the suffix of the created sql files.
```shell
make migrate
```
Execute migratedown command
```shell
make migratedown
```
Execute migrateup command
```shell
make migrateup
```

