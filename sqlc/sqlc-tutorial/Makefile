sqlite3def := go run github.com/sqldef/sqldef/cmd/sqlite3def@v0.17.6
sqlc := go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0

.PHONY: migrate
migrate:
	$(sqlite3def) -f schema.sql db.sqlite

.PHONY: migrate-dry
migrate-dry:
	$(sqlite3def) -f schema.sql --dry-run db.sqlite

.PHONY: generate
generate:
	$(sqlc) generate

