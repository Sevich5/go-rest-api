# Dockerized Golang REST Api boilerplate: nginx + gin + gorm + PostgreSQL(with atlas migrations) + JWT Auth  

### Install:
**PostgreSQL must be 15 version**

````
go install ariga.io/atlas/cmd/atlas@latest
go mod download
````

### Working with migrations
**It must be second database for creating migrations with atlas**
#### Make diff migration (creates sql file)
````
make db-diff
````
#### Apply migration
````
make db-migrate
````
#### Migrations status
````
make db-migration-status
````

## Creating user from cli
````
go run cmd/api/create_user.go [email] [password]
````
## Running server
````
go run cmd/api/main.go
````
