data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/api/migrate.go",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  migration {
    dir = "file://internal/infrastructure/persistence/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}