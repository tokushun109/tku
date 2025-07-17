---
allowed-tools: Bash(git add:*), Bash(git commit:*), Bash(git push:*), Bash(gh *)
description: "Issue 番号を受け取り、修正 → commit → push → PR 作成まで完走します（マージはしません）"
---

## Context

- Issue 番号：$ARGUMENTS

## Your task

1. `gh issue view $ARGUMENTS` で内容を取得・読み込み
2. 修正を実行（プロジェクト状況に応じた変更）
3. `git add .` → コミットメッセージを作成して `git commit -m "Issue #$ARGUMENTS: <summary>"`
4. コミットにエラーがあれば自動修正・再コミット
5. `git push -u origin $(git branch --show-current)`
6. プッシュ後のテストでエラー発生 → 修正 → 再コミット・プッシュ
7. `gh pr create --title "Fixes #$ARGUMENTS: <title>" --body "<修正内容>"` で PR 作成
8. PR 本文に修正内容を記述
9. 自分でレビューを行い、「承認理由」と「次のアクション（例：待機 or 次対応者指定）」を本文に追加
10. **マージは絶対に行わない**
