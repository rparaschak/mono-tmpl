---
name: backend-testing
description: Backend testing guidance and checklist.
---

# Backend Testing

## Strategy (EXTREMELY IMPORTANT)
- Primary testing through API
- Cases that can not be tested through API are tested through usecase tests
- No need to test every single edge case - focus on the most important ones
- Tests are readable.
- Test setup and data creation should be encapsulated 

## Comon Rules
- Build tag: `//go:build integration`
- Max Depth = 3 
- API Test Structure: `TestFunction` → `t.Run("[METHOD] /endpoint")` → `t.Run("Behavior")`
- Usecase Test Structure: `TestFunction` → `t.Run("Feature Group")` → `t.Run("Behavior")`
