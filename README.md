## install

```shell
cp -p .envrc.sample .envrc
# for sqlc generates fully type-safe idiomatic Go code from SQL. 
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# for debug
go install github.com/go-delve/delve/cmd/dlv@latest
# for hot reload
go install github.com/cosmtrek/air@latest
# download module
go mod download
air
```

## update query.sql.go

1. change `./sqlc/schema.sql`, `./sqlc/query.sql`
2. run sqlc
 
    ```shell
    sqlc generate --experimental
    ```

3. generate `./sqls/models.go`, `./sqls/query.sql.go`
 