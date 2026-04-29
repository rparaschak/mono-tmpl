---
name: backend-code-review
description: Backend code review guidance and checklist.
---

# Backend Code Review

## Checklist

### Handlers
- [ ] Handler request payloads should reuse shared `contracts` DTOs where possible (avoid duplicating body field structs inside `handlers/*`)
- [ ] Handlers should be thin and delegate to usecases
- [ ] Handlers should not contain business logic

### Query Performance
- [ ] No N+1 queries - prefer batch operations
- [ ] Never use JOIN/PRELOAD on models from another module
- [ ] No queries inside loops
- [ ] Batch operations for multiple records

### Code Simplification
- [ ] No defensive coding (no search before delete/update)

### Tests
- [ ] Split files: `{feature}_usecase_test.go` + `{feature}_api_test.go`
- [ ] API tests don't duplicate usecase test behaviors
- [ ] Access control tests consolidated (one per role)
