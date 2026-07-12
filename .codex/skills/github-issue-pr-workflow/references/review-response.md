# Review Response Guidance

Use this reference only after PR review feedback exists.

## Triage

- Separate blocking requested changes from optional suggestions.
- Group duplicate comments before editing.
- Prefer code changes when the reviewer points out a concrete defect, failing test, unclear API, accessibility issue, security risk, or maintainability problem.
- Prefer explanation when the comment is based on a mistaken assumption and the current code is correct. Include evidence from code, tests, docs, or product requirements.
- Ask a follow-up question when the requested behavior conflicts with the issue scope or would introduce a larger design decision.

## Response Plan Format

Keep the plan short:

```markdown
対応方針:
- <comment summary>: 修正します。
- <comment summary>: 既存仕様と照合し、PR上で説明します。
- <comment summary>: 追加確認が必要なため質問します。
```

## Applying Fixes

- Read the surrounding code before changing anything.
- Preserve unrelated user changes in the worktree.
- Keep each commit focused on the review response.
- Add or adjust tests when the review identifies behavior that should not regress.
- Re-run the smallest reliable validation set, then broaden it if shared behavior changed.

## Replying

- Be concise and specific.
- Mention the commit or change that addresses the comment when possible.
- If not changing code, explain why and reference the relevant requirement or existing behavior.
- Do not mark a thread resolved unless the requested action is complete or the reviewer has accepted the explanation.
