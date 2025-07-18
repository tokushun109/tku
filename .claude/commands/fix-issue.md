---
allowed-tools: Task, Bash(git add:*), Bash(gh *)
description: "Issue 番号を受け取り、修正 → commit → push → PR 作成まで完走します（マージはしません）"
---

## Context

- Issue 番号：$ARGUMENTS

## Your task

**必ずを下記のステップで、一つも省略せずに実行すること**

1. `gh issue view $ARGUMENTS` で内容を取得・読み込み
2. issue 番号に合わせてブランチを作成(feature/#$ARGUMENTS)
3. 修正を実行（プロジェクト状況に応じた変更）
4. `git add .` → Claude Code の組み込み機能を使用してコミット・プッシュ
5. コミットにエラーが出た時は`pnpm lint-fix`を実行して再コミットを行い、それでもダメならエラーに応じた修正を実施して再コミットを行う
6. Claude Code の組み込み機能を使用してプッシュ(**git commit は使わずに claude -p で実行**")
7. プッシュ後のテストでエラー発生 → 修正 → 再コミット・プッシュ(**git push は使わずに claude -p で実行**")
8. Claude Code の組み込み機能を使用して PR 作成
9. PR 本文に修正内容を記述
10. 自分でレビューを行い、「承認理由」と「次のアクション（例：待機 or 次対応者指定）」を本文に追加
11. **マージは絶対に行わない**

## 実装方法

### 手順 3-7: Claude Code の組み込み機能を使用

```bash
# Task toolを使用してコミット・プッシュ・PR作成を実行
Task:
  description: "Git operations and PR creation"
  prompt: |
    現在ステージングされている変更を以下の手順で処理してください：

    1. コミットメッセージ「Issue #$ARGUMENTS: <summary>」でコミット作成
    2. 現在のブランチにプッシュ
    3. タイトル「Fixes #$ARGUMENTS: <title>」でPR作成
    4. PR本文に修正内容を詳細に記述
    5. レビューを実施し、承認理由と次のアクションを追加
    6. マージは絶対に行わない
```

### 注意事項

- `git commit`と`git push`は権限制限のため直接実行できません
- 代わりに Task tool を使用して Claude Code の組み込み機能を活用します
- すべての git 操作は Task tool 経由で実行してください
