---
name: new-project
description: Initialize this monorepo template for a new project. Use only when the user explicitly asks for `new-project` and provides a project name. Generates infrastructure files from `.tmpl` files and replaces `{XXX}` placeholders with the given project name.
---

# New Project

Use this skill only on explicit user request.

## Input

- Expect exactly one required argument: the project name.
- Accept lowercase letters, digits, and hyphens only: `^[a-z0-9][a-z0-9-]*$`

## Workflow

1. Validate that the user provided a project name.
2. Run `bash .agents/skills/new-project/scripts/new_project.sh <project-name>` from the repository root.
3. Report which infrastructure files were generated or changed.

## Notes

- This skill currently updates placeholders only under `infrastructure/`.
- `*.tmpl` files are source templates and must not be modified.
- Do not guess a project name. If the user did not provide one, ask for it.
