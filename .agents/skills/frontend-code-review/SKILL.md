---
name: frontend-code-review
description: Frontend code review guidance and checklist.
---

# Frontend Code Review

## Checklist

### Architecture
- [ ] Frontend code follows the flows and contracts defined in the plan.
- [ ] Business rules are not duplicated in presentation code when they belong in backend or shared layers.

### Components
- [ ] Components are focused and avoid mixing data fetching, state orchestration, and rendering without a reason.
- [ ] Reused UI patterns stay consistent with the existing frontend structure.

### State And Data
- [ ] Async states are handled explicitly: loading, success, empty, and error.
- [ ] Backend contracts, field names, and nullability are handled correctly.
