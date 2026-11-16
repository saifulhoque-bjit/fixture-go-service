# fixture-go-service

A small Go HTTP service that keeps a user directory in SQLite (pure Go driver,
`modernc.org/sqlite`, so no C compiler is needed).

This is a BJIT test fixture for AMS. It is not a product.

    go run .                 # serves on 127.0.0.1:8080
    curl 'http://127.0.0.1:8080/users?name=alice'
    go test ./...
