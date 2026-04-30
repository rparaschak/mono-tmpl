---
name: backend-architecture
description: Backend architecture and implementation conventions for module structure, usecase boundaries, handlers, events, jobs, MCP tools, error handling, and testing.
---

# Backend Architecture Reference

## Architecture

```text
Request <-(DTO)-> Handler|MCP <-(Domain)-> UseCase <-(Domain)-> Persistence
Trigger(Job,Event) -> (Domain)-> UseCase <-(Domain)-> Persistence
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
- Do not import `contracts/` into usecases; map transport DTOs before the usecase boundary.
- Prefer explicit `Input` structs for command-style usecases over accepting `models.Model`.
- Do not create `Input` structs that only wrap a single primitive value or ID; accept that value directly, e.g. `DeleteSample(ctx, sampleId uuid.UUID)`.
- Accept `models.Model` only when the whole domain object is truly the business input, not just a convenient field container.
- Prefer batch usecases over Preload/Join; compose separate usecases instead.
- Never use JOIN/PRELOAD on models from another module; use batch instead.
- No database queries inside loops; use `WHERE id IN (?)` patterns.
- No protective coding; do not search before deleting/updating.
- Never use `fmt.Errorf`; return predefined errors from `errors.go`.

## Error Handling

- Modules define errors in `[module]/errors/errors.go` as `&coreErrors.AppError{Status: http.StatusNotFound, Message: "..."}`.
- `pkg/` packages use `errors.New()`.

## REST Handlers

- File naming: `[action]_[resource].go` (one handler per file).
- Naming: `[Action][Entity]Handler`, `[Action][Entity]Request`, `[Action][Entity]Response`.
- Params: `camelCase`, `uuid.UUID` for IDs, struct tags (`format:"uuid"`, `minLength:"1"`, `enum:"..."`).
- Body fields are required by default; use `omitempty` or `omitzero` for optional fields.
- `required` tag is only for header/query params (optional by default).
- No business logic in handlers; delegate to usecases.
- Registration pattern: `routing.POST(groups.Auth, "", "Create Sample", h.CreateSampleHandler)`.
- Keep read DTOs response-only. Request bodies must bind dedicated input DTOs from `contracts/`, not response DTOs with `readOnly` fields.
- Reuse shared `*InputDTO` contracts across create/update handlers when the writable shape is the same; introduce operation-specific input DTOs only when the write shapes differ.

## Handler/Usecase Communication

- Keep the mapping at the boundary: handlers map `contracts/*InputDTO` into `usecases/*Input` or domain values before calling a usecase.
- Handlers must not pass `contracts` DTO types into usecases.
- `contracts/*InputDTO` is transport-only and may be reused across handlers when the writable HTTP payload shape is the same.
- If create and update belong to the same command family, prefer the same signature style for both, ideally explicit usecase input structs.
- If a model includes DB-managed fields like `Id`, `CreatedAt`, or `UpdatedAt`, do not use that model as the handler-to-usecase input shape.
- Using `models.Model` in a trivial create usecase is an allowed shortcut, but it is an exception and not the default rule.
- Once a usecase gains validation, orchestration, derived fields, permissions, or events, replace `models.Model` input with an explicit `usecases.Input`.

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
