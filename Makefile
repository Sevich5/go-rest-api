DSN_DEV_VAR := $$(go run cmd/api/migrate.go dsn_dev)
DSN_VAR := $$(go run cmd/api/migrate.go dsn)

build:
	go build -o build/api cmd/api/main.go
run:
	go run cmd/api/main.go
db-diff:
	atlas migrate diff --env gorm --dev-url="$(DSN_DEV_VAR)"
db-migrate:
	atlas migrate apply --env gorm --url="$(DSN_VAR)"
db-migration-new:
	atlas migrate new
db-migration-status:
	atlas migrate status --env gorm --url="$(DSN_VAR)"