## install

make setup

## update query.sql.go

1. change `./sqlc/schema.sql`, `./sqlc/query.sql`
2. run sqlc
 
    ```shell
    make sqlcg
    ```

3. generate `./sqls/models.go`, `./sqls/query.sql.go`
 
