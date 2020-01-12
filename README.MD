## tldr-web backend

### before use

use `database/create_table.sql` create database and data table.

```shell script
export GITHUB_USER='<github_username>'
export GITHUB_PASS='<github_password>'
```

### usage

```shell script
go run main.go
```

will automatic sync data from GitHub into database.
