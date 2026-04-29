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

targets=()

while IFS= read -r template; do
  output="${template%.tmpl}"
  perl -0pe "s/\\{XXX\\}/$project_name/g" "$template" > "$output"
  targets+=("$output")
done < <(find "$infrastructure_dir" -type f -name '*.tmpl' | sort)

while IFS= read -r file; do
  perl -0pi -e "s/\\{XXX\\}/$project_name/g" "$file"
  targets+=("$file")
done < <(rg -l --glob '!*.tmpl' '\{XXX\}' "$infrastructure_dir")

if [[ ${#targets[@]} -eq 0 ]]; then
  echo "No infrastructure templates or {XXX} placeholders found under $infrastructure_dir" >&2
  exit 1
fi

for target in "${targets[@]}"; do
  echo "$target"
done
