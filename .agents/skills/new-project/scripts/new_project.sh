#!/usr/bin/env bash

set -euo pipefail

if [[ $# -lt 1 || $# -gt 2 ]]; then
  echo "Usage: $0 <project-name> [repo-root]" >&2
  exit 1
fi

project_name="$1"
repo_root="${2:-.}"
infrastructure_dir="${repo_root%/}/infrastructure"

if [[ ! "$project_name" =~ ^[a-z0-9][a-z0-9-]*$ ]]; then
  echo "Invalid project name: $project_name" >&2
  echo "Expected lowercase letters, digits, and hyphens only." >&2
  exit 1
fi

if [[ ! -d "$infrastructure_dir" ]]; then
  echo "Infrastructure directory not found: $infrastructure_dir" >&2
  exit 1
fi

files=()
while IFS= read -r file; do
  files+=("$file")
done < <(rg -l '\{XXX\}' "$infrastructure_dir")

if [[ ${#files[@]} -eq 0 ]]; then
  echo "No {XXX} placeholders found under $infrastructure_dir" >&2
  exit 1
fi

for file in "${files[@]}"; do
  perl -0pi -e "s/\\{XXX\\}/$project_name/g" "$file"
  echo "$file"
done
