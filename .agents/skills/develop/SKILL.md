---
name: develop
description: Orchestrate backend and frontend implementation, testing, and code review with specialized agents.
---

# Develop
Orchestration skill. You are not allowed to write code yourself. You can only delegate to other agents.
Input comes from a feature folder: `docs/development/<x-feature-name>/` with `plan.md` and `spec.md`.
If the user does not specify a feature folder, identify the most recent folder candidate and ask the user to confirm before proceeding.

## WORKFLOW

1. Confirm feature folder and documentation paths (`plan.md` and `spec.md`).
2. Backend implementation. Delegate to `backend-engineer`.
3. Frontend implementation. Delegate to `frontend-engineer` after backend implementation is complete.
4. Test implementation. Delegate to `backend-test-engineer`.
5. Verify. Run `make check`.
6. Review. If backend changes exist, delegate to `backend-code-reviewer`. If frontend changes exist, delegate to `frontend-code-reviewer`. Run both reviews in parallel when both apply.
7. Fix. If backend issues exist, delegate to `backend-engineer`. If frontend issues exist, delegate to `frontend-engineer`. Run both fix tasks in parallel when both apply.


# Delegation rules

1. Provide path to product document and spec to every agent.
2. Paths must point to `docs/development/<x-feature-name>/plan.md` and `docs/development/<x-feature-name>/spec.md`.
3. If schema or model changes are required, instruct backend engineer to use `migrations` skill.
4. Backend work should be completed before delegating frontend work.
5. Backend engineer owns backend implementation; frontend engineer owns frontend implementation.
6. Backend code review is owned by `backend-code-reviewer`; frontend code review is owned by `frontend-code-reviewer`.
7. Run code review only for layers that changed.
8. When both backend and frontend changed, run both code reviews in parallel.
9. Run fixes only for layers with findings or verification failures.
10. When both backend and frontend need fixes, run both fix tasks in parallel.
11. If you identify multiple feature slices, suggest a plan to the user how to split them.

## Tests
- Tests should be based on behaviors given in the task. When behaviors are not given, create a list of behaviors so I can verify and approve.
