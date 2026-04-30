---
name: backend-architecture
description: Backend architecture and implementation conventions for module structure, usecase boundaries, handlers, events, jobs, MCP tools, error handling, and testing.
---

# Backend Architecture Reference

## Architecture

```text
Request <-(DTO)-> Handler|MCP <-(UsecaseInput|Primitive)-> UseCase <-(Domain)-> Persistence
Trigger(Job,Event) -> (UsecaseInput|Primitive)-> UseCase <-(Domain)-> Persistence
```

## Module Structure

```text
modules/[module]/
├── router.go           # Route definitions
├── handlers/           # HTTP handlers
├── usecases/           # Business logic
├── models/             # Domain models
├── contracts/          # DTOs
├── events/             # Event handlers
├── jobs/               # Background jobs
├── mcp/                # AI assistant tools
├── testing/            # Test factories and data
└── errors/             # Business errors
```

- Modules have dependencies on each other.
- Circular dependencies are not allowed.

## Usecases

### Philosophy

- Simple: each usecase does one thing well.
- Reusable: callable from handlers, jobs, event handlers, or other usecases.
- Composable: complex workflows should compose multiple usecases.
- Foundation: jobs, event handlers, and HTTP handlers are based on usecases.

### Rules

- Use GORM generics: `gorm.G[Model](db)`.
- Keep usecase inputs transport-agnostic; do not add `json`, schema, or HTTP-specific tags to `usecases/` types.
- Take Input structs or primitives for simple inputs, e.g. `DeleteSample(ctx, sampleId uuid.UUID)`.
- Prefer batch(e.g. GetSamplesByIdsBatch) usecases over Preload/Join; compose separate usecases instead.
- Batch usecases return maps, e.g. `map[uuid.UUID]Model{}`.
- Usecases might share Input models defined in `usecases.go`
- No database queries inside loops.
- No protective coding; do not search before deleting/updating.
- Never use `fmt.Errorf`; return predefined errors from `errors.go`.

## Error Handling

- Modules define errors in `[module]/errors/errors.go` with `pkg/apperror`.
- Keep `apperror.New(...)` declarations on one line, e.g. `var ErrSampleNotFound = apperror.New(http.StatusNotFound, "sample_not_found", "sample not found")`.
- Use stable snake_case error codes. Error messages are short, lowercase, and safe to return to clients.
- Use `apperror.WithDetails(err, details)` or `err.WithDetails(details)` for per-request details; do not mutate package-level errors.
- Usecases translate known infrastructure errors into predefined module errors.
- Handlers return usecase errors directly. Huma uses `GetStatus()` from `apperror.AppError` to set the HTTP status.
- `pkg/` packages use `errors.New()`.

## API Handlers

- No business logic in handlers; delegate to usecases.
- Body fields are required by default; use `omitempty` or `omitzero` for optional fields.
- `required` tag is only for header/query params (optional by default).
- Reuse shared `*InputDTO` contracts across create/update handlers when the writable shape is the same; introduce operation-specific input DTOs only when the write shapes differ.
- Reuse `*DTO` for read operation if possible.
- Never pass `contracts` DTO types into usecases.

## Events

- Topics and payload structs are defined in `modules/core/events.go`.
- Topic naming: `ModuleName.EventName`.
- Publish with `u.Events.Publish(core.TopicName, PayloadStruct{...})`.
- Subscribe with `events.On(s, core.TopicName, "HandlerName", e.Handler)` in `Register()`.
- Event handlers are slim; no business logic, always call usecase.
- Handler errors are automatically logged.

## Jobs

- Job types and payload structs are defined in `modules/core/jobs.go`.
- Register with `j.ScheduledJobs.Subscribe(core.JobType, handler)` in `Register()`.
- Schedule from usecases only: `u.ScheduledJobs.Schedule(ctx, jobType, &payload, startAfter)`.
- Job handlers unmarshal payload and call usecases; no business logic.
- Flow: `UseCase.Schedule() -> DB Queue -> Subscriber -> Handler -> UseCase.Execute()`.

## MCP Handlers

- MCP tools live in `modules/[module]/mcp/` only.
- Define input/output structs with `json`, `jsonschema_description`, and `jsonschema` tags.
- Register tools in the module `WithMCP()` method.
- Delegate to usecases; no business logic in MCP handlers.
- All errors are defined in the module `errors/errors.go`.

## Testing

- Primary: API tests.
- Secondary: usecase tests.
- Everything not covered by API tests should be covered by usecase tests.
