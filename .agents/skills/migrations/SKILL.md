---
name: migrations
description: Working rules and commands for schema migrations in the API project.
---

# Migrations

Use this skill whenever you change database schema or add/update models.

## Rules

1. New models must be added to `api/cmd/migrator/main.go`.
2. Do not edit `api/migrations/atlas.sum` manually.
3. Run commands from the `api/` directory.

## Commands

### Migrate models (schema diff)
```bash
make migration NAME=feature_name ENV=local
```

### Create a manual migration
```bash
atlas migrate new feature_name --env local
```

Implement the generated migration file, then:

```bash
atlas migrate hash --env local
```

Equivalent Make targets:
```bash
make migration-new NAME=feature_name ENV=local
make migration-hash ENV=local
```

## Verification

1. Ensure migration files are generated in `api/migrations/`.
2. Ensure model changes are reflected in `api/cmd/migrator/main.go`.
