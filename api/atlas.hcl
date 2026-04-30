data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/migrator",
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  dev = "docker+postgres://imresamu/postgis:17-3.4-alpine/dev?search_path=public"
  url = "postgresql://supabase_admin:docker@localhost:5002/mono-tmpl?search_path=public&sslmode=disable"

  migration {
    dir = "file://migrations"
  }
}

env "production" {
  src = data.external_schema.gorm.url
  url = getenv("DATABASE_URL")

  migration {
    dir = "file://migrations"
  }
}
