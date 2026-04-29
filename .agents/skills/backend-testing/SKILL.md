---
name: backend-testing
description: Backend testing guidance and checklist.
---

# Backend Testing

## Strategy (EXTREMELY IMPORTANT)
- Primary testing through API
- Cases that can not be tested through API are tested through usecase tests
- No need to test every single edge case - focus on the most important ones
- Tests are readable. Test setup and data creation should be encapsulated 

## API Tests

### File Structure
- Filename: `{handler_name}_test.go`
- Location: same directory as handler file
- Build tag: `//go:build integration`
- One file per feature

### Naming: `func [FeatureName]_API_Test`
```go
func BrandInvitations_API_Test(t *testing.T) {
    t.Run("POST /brands/:brandId/invitation", func(t *testing.T) {
        t.Run("Forbidden for anonymous", func(t *testing.T) {...})
        t.Run("Forbidden for non-manager", func(t *testing.T) {...})
    })
    t.Run("POST /invitations/:token/accept", func(t *testing.T) {
        t.Run("Forbidden for anonymous", func(t *testing.T) {...})
    })
}
```

### Max Depth = 3
`TestFunction` → `t.Run("[METHOD] /endpoint")` → `t.Run("behavior")`

## Usecase Testing

### File Structure
- Filename: `{usecase_name}_test.go`
- Location: same directory as usecase file
- Build tag: `//go:build integration`
- One file per feature

### Naming: `func [FeatureName]_Usecase_Test`
```go
func BrandInvitations_Usecase_Test(t *testing.T) {
    t.Run("Sending an invite", func(t *testing.T) {
        t.Run("User can send invite", func(t *testing.T) {...})
        t.Run("Rejects for existing member", func(t *testing.T) {...})
    })
    t.Run("Accepting an invite", func(t *testing.T) {
        t.Run("Successfully accepts the invite", func(t *testing.T) {...})
    })
}
```

### Max Depth = 3
`TestFunction` → `t.Run("Group")` → `t.Run("behavior")`
