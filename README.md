# SQLite Scheme Migration Bug

This is a repository to reproduce a bug with Ent and the [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) non-CGO SQLite 3 driver.

## Reproduce

Simply run `go test ./...`. The test will run `Schema.Create()` twice from the generated Ent client. In the real world this is reproducible by starting
the application twice when using a SQLite file-based DB. For this bug report I am simply running it twice in the test.

This works fine

## Expected Results

I expect that the second pass makes no changes to the database and therefore there should be no failures.