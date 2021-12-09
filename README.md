# SQLite Schema Migration Bug

This is a repository to reproduce a bug with Ent and the [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) (non-CGO) SQLite 3 driver.

## Reproduce

Simply run `go test ./...`. The test will run `Schema.Create()` twice from the generated Ent client. In the real world this is reproducible by starting
the application again when using a SQLite file-based DB that already exists. For this bug report I am simply running it twice in the test.

The test shows how this works fine with `github.com/mattn/go-sqlite3`, but the CGO requirement is a non-starter for my usecase and using the CGO-free driver is the only real option.

## Expected Results

I expect that the second pass returns no errors, exactly like `mattn/go-sqlite3` does.
