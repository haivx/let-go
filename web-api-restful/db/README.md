Using golang-migrate for running makefile

1. Command create schema

```
  migrate create -ext sql -dir  db/migration -seq init_schema
```

2. Run migrate

```
  migrate -path db/migration -database "postgresql://root:secret@localhost:5432/final-project?sslmode=disable" -verbose up
```