---
name: github-issue-pr-workflow
description: Use this skill when the user asks Codex to handle a repository task end-to-end through GitHub: create an issue from the request, create or switch to a branch, implement the change, open a pull request with an appropriate Japanese description, wait for or fetch PR review feedback, decide a response plan, apply fixes, and push updates.
---

# GitHub Issue PR Workflow

## Overview

This skill runs a full GitHub-driven development loop for this repository: request intake, issue creation, branch work, implementation, PR creation, review handling, and follow-up pushes.

Use it when the user wants work managed through GitHub rather than only local edits.

## Required Context

- Read `AGENTS.md` and the relevant `CLAUDE.md` files before editing.
- Inspect `git status --short` before changes. Do not overwrite unrelated user changes.
- Prefer GitHub MCP tools when available. Otherwise use `gh` CLI. If neither is authenticated, stop after local preparation and explain the blocker.
- Use Japanese for user-facing updates, issue bodies, PR descriptions, and commit messages unless the repository convention says otherwise.

## Workflow

### 1. Intake and Issue

1. Convert the user's request into a concise issue title and body.
2. Search existing open issues first to avoid duplicates.
3. Create a GitHub issue only when no suitable issue already exists.
4. Include:
   - Background or goal
   - Requested behavior
   - Scope and exclusions when known
   - Validation expectations

If the request is ambiguous enough that implementation would be risky, ask one concise clarification before creating the issue.

### 2. Branch

1. Sync the default branch using the repository's normal flow.
2. Create a branch named from the issue, for example `issue-123-fix-product-image-upload`.
3. If an appropriate branch already exists, switch to it after checking it is not carrying unrelated work.

Do not use destructive git commands. Avoid force-push unless the user explicitly approves it.

### 3. Implement

1. Make the smallest coherent change that satisfies the issue.
2. Follow the repository guides:
   - Root: `AGENTS.md`, `CLAUDE.md`
   - Frontend: `frontend/CLAUDE.md`
   - Backend: `backend/CLAUDE.md`
   - Infra: `infra/CLAUDE.md`
3. Keep commits focused. Do not include `.gitignore` targets, secrets, build artifacts, or local editor settings.
4. Run targeted tests, lint, build, or type checks proportional to the change.

When a change spans frontend/backend/infra, validate each touched area enough to catch contract errors.

### 4. Pull Request

1. Push the branch.
2. Create a PR linked to the issue.
3. Write the PR body in the repository's expected style. If no template exists, use:

```markdown
## 目的/背景

## 変更点（要点箇条書き）

## 影響範囲（UI/SEO/DB/インフラ）

## 検証手順（実行コマンドと観点）

## リスクとロールバック

## 参照資料（関連する CLAUDE.md 節、Issue など）

Closes #<issue-number>
```

4. Include exact validation commands and results. If a check could not be run, say why.

### 5. Review Handling

When the user says reviews are available, or asks Codex to continue after review:

1. Fetch PR review comments, requested changes, unresolved threads, and CI status.
2. Read each comment in context before editing.
3. Decide and state a short response plan:
   - Fix now
   - Explain or push back with evidence
   - Ask the reviewer for clarification
4. Apply accepted fixes locally.
5. Re-run relevant validation.
6. Commit and push the updates.
7. Reply to resolved review threads when the tool surface supports it.

For detailed triage guidance, read `.codex/skills/github-issue-pr-workflow/references/review-response.md` only when review feedback exists.

## Tooling Patterns

With GitHub MCP, use repository-aware tools for issue search/creation, branch/PR operations, review reading, and review replies.

With `gh` CLI, typical commands are:

```bash
gh issue list --search "<keywords> state:open"
gh issue create --title "<title>" --body "<body>"
git switch -c issue-123-short-slug
git push -u origin issue-123-short-slug
gh pr create --fill --body-file <file>
gh pr view --comments
gh pr checks
```

Keep temporary PR body files outside committed changes or remove them before finishing.

## Done Criteria

- Issue exists and is linked from the PR.
- Branch is pushed.
- Implementation is complete and scoped.
- PR description explains purpose, changes, impact, validation, and risk.
- Review feedback, if present, is either addressed, answered, or explicitly blocked on clarification.
- Final user update includes issue/PR numbers or URLs, validation results, and any remaining risks.
