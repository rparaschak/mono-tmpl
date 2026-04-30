---
name: backend-testing
description: Backend testing guidance and checklist.
---

# Backend Testing

## Strategy (EXTREMELY IMPORTANT)
- Primary testing through API.
- API tests use the real router, real handlers, real usecases, and real infrastructure.
- API tests use `github.com/gavv/httpexpect/v2` bound to the in-process router.
- Cases that can not be tested through API are tested through usecase tests.
- Usecase tests are also appropriate when HTTP setup would obscure the behavior under test.
- No need to test every single edge case - focus on the most important ones.
- Test setup and data creation should be encapsulated.
- Create domain data through usecases by default.
- Do not seed API tests by calling HTTP endpoints, except when testing the create endpoint itself.
- Direct DB writes are allowed only inside test helpers for fields no usecase can express, such as controlled timestamps.
- Test data names should be unique enough to avoid matching migration seed rows or data created by other tests.

## Common Rules
- Build tag: `//go:build integration`
- File names: handler tests end with `_api_test.go` and usecase tests end with `_usecase_test.go`
- Max Depth = 3 
- API test files are named by endpoint behavior, for example `create_sample_api_test.go`.
- API Test Structure: `TestFunction` → `t.Run("METHOD /endpoint")` → `t.Run("behavior")`
- Usecase Test Structure: `TestFunction` → `t.Run("Feature Group")` → `t.Run("Behavior")`
- Use `github.com/stretchr/testify/require` for setup and fatal assertions.
- Use `github.com/stretchr/testify/assert` for readable non-fatal checks.
- `require` and `assert` calls should include explanation strings.
- Shared integration helpers live in `api/internal/testsupport/integration`.
- Module-specific fixtures, predefined DTOs, and helper assertions live in a module-local `testkit` package.
- Do not use `t.Parallel()` for tests that share the real integration database unless the data is isolated per test.
- Run shared-DB integration packages sequentially with `go test -p 1`.
