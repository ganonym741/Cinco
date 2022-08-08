# CINCO
### CREATE DATABASEUSING POSTGRES
> Create database "**cinco**"

create file **.env** then fill these to configure Makefile:
>DB_USER=postgres
DB_PASSWORD=password
DB_PORT=5432
DB_NAME=cinco


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

